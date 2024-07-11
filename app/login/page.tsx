"use client";
import React from "react";

function LoginPage() {
  function onSubmit(event) {
    event.preventDefault();
    console.log("salam");
    try {
      const formData = new FormData(event.currentTarget);
      for (let [key, value] of formData.entries()) {
        console.log(key, value);
      }
    } catch (error) {
      console.log("ridi");
    }
  }

  return (
    <div className="flex flex-col w-fit mx-auto h-screen gap-1 justify-center items-center">
      <form onSubmit={onSubmit}>
        <div className="flex flex-col gap-2 justify-center items-center">
          <input
            type="text"
            name="username"
            className="border-2 border-green-400"
            placeholder="نام کاربری"
          />
          <input
            type="password"
            name="password"
            className="border-2 border-green-400"
            placeholder="رمز عبور"
          />
        </div>
        <div className="bg-green-700 p-2 rounded-xl text-center text-white">
          <button type="submit">بزن بریم!</button>
        </div>
      </form>
    </div>
  );
}

export default LoginPage;
