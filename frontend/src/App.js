import Login from "./components/login";
import SignUp from "./components/SignUp"
import "bootstrap/dist/css/bootstrap.min.css";
import "./style.css";
import { Route, Routes } from "react-router-dom";
import Home from "./pages/Home";
import Navbar from "./components/Navbar";
import Dashboard from "./pages/Dashboard";
import CreateCoffee from "./pages/CreateCoffee";

export default function App() {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/Login" element={<Login />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/farmer/createCoffee" element={<CreateCoffee />} />
      </Routes>
    </>
  );
}