"use client";
import React, { useState, useEffect } from "react";

interface Unit {
  id: number;
  name: string;
  description: string;
}

const UnitPage: React.FC = () => {
  const [units, setUnits] = useState<Unit[]>([]);
  const [newUnit, setNewUnit] = useState({ name: "", description: "" });

  useEffect(() => {
    // Fetch units from the server when the component mounts
    fetchUnits();
  }, []);

  const fetchUnits = async () => {
    try {
      const response = await fetch("http://localhost:3000/units");
      const data = await response.json();
      setUnits(data);
    } catch (error) {
      console.error("Error fetching units:", error);
    }
  };

  const addUnit = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      const response = await fetch("http://localhost:3000/units", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newUnit),
      });
      const data = await response.json();
      setUnits([...units, data]);
      setNewUnit({ name: "", description: "" });
    } catch (error) {
      console.error("Error adding unit:", error);
    }
  };

  const deleteUnits = async (event: React.FormEvent) => {
    event.preventDefault();
    const checkedUnits = units.filter(
      (unit) =>
        (document.getElementById(`unit-${unit.id}`) as HTMLInputElement).checked
    );
    try {
      await Promise.all(
        checkedUnits.map((unit) =>
          fetch(`http://localhost:3000/units/${unit.id}`, {
            method: "DELETE",
          })
        )
      );
      setUnits(units.filter((unit) => !checkedUnits.includes(unit)));
    } catch (error) {
      console.error("Error deleting units:", error);
    }
  };

  return (
    <div className="h-screen flex flex-col gap-4 p-4">
      <h1 className="text-xl font-bold">Unit Page</h1>

      <form onSubmit={addUnit} className="flex flex-col gap-2">
        <input
          type="text"
          name="name"
          value={newUnit.name}
          onChange={(e) => setNewUnit({ ...newUnit, name: e.target.value })}
          placeholder="Unit Name"
          className="border p-2"
          required
        />
        <input
          type="text"
          name="description"
          value={newUnit.description}
          onChange={(e) =>
            setNewUnit({ ...newUnit, description: e.target.value })
          }
          placeholder="Unit Description"
          className="border p-2"
          required
        />
        <button type="submit" className="bg-blue-500 text-white p-2 rounded">
          Add Unit
        </button>
      </form>

      <form onSubmit={deleteUnits} className="flex flex-col gap-2">
        <div className="flex flex-col gap-2">
          {units.map((unit) => (
            <label key={unit.id} className="flex items-center gap-2">
              <input type="checkbox" id={`unit-${unit.id}`} />
              {unit.name} - {unit.description}
            </label>
          ))}
        </div>
        <button type="submit" className="bg-red-500 text-white p-2 rounded">
          Delete Selected Units
        </button>
      </form>
    </div>
  );
};

export default UnitPage;
