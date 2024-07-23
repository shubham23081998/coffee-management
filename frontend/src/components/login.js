import React, { useState } from "react";
import { Form, Button, Alert } from "react-bootstrap";
import "./login.css";
import { Link, useNavigate } from "react-router-dom";

import BackgroundImage from "../assests/background.png";
import Logo from "../assests/signInLogo.jpg";
import { login } from "../restService";

const Login = () => {
  const [inputEmail, setInputEmail] = useState("");
  const [inputPassword, setInputPassword] = useState("");
  const [inputRole, setInputRole] = useState("");
  const navigate = useNavigate();

  const [show, setShow] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);
    await delay(500);
    login(inputEmail, inputPassword, inputRole).then((resp) => {
      console.log(resp);
      if (resp.data != null) {
        localStorage.setItem("user", JSON.stringify(resp.data.data));
        navigate("/dashboard")
      }
    })
    setLoading(false);
  };

  function delay(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  const handleSignUpClick = () => {
    navigate("/signup")
  };

  const handleOnChangeROle = (event) => {
    setInputRole(event.target.value);
  }

  return (
    <div
      className="sign-in__wrapper"
      style={{ backgroundImage: `url(${BackgroundImage})` }}
    >
      {/* Overlay */}
      <div className="sign-in__backdrop"></div>
      {/* Form */}
      <Form className="shadow p-4 bg-white rounded" onSubmit={handleSubmit}>
        {/* Header */}
        <img
          className="img-thumbnail mx-auto d-block mb-2"
          src={Logo}
          alt="logo"
        />
        <div className="h4 mb-2 text-center">Sign In</div>
        {/* ALert */}
        {show ? (
          <Alert
            className="mb-2"
            variant="danger"
            onClose={() => setShow(false)}
            dismissible
          >
            Incorrect Email or password.
          </Alert>
        ) : (
          <div />
        )}
        <Form.Group className="mb-2" controlId="email">
          <Form.Label>Email</Form.Label>
          <Form.Control
            type="email"
            value={inputEmail}
            placeholder="Email"
            onChange={(e) => setInputEmail(e.target.value)}
            required
          />
        </Form.Group>
        <Form.Group className="mb-2" controlId="password">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            value={inputPassword}
            placeholder="Password"
            onChange={(e) => setInputPassword(e.target.value)}
            required
          />
        </Form.Group>
        <Form.Group className="mb-3">
          <Form.Label htmlFor="disabledSelect">Role</Form.Label>
          <Form.Select onChange={handleOnChangeROle}>
            <option value={""} disabled selected>Select</option>
            <option value={"farmer"}>Farmer</option>
            <option value={"processor"}>Processor</option>
            <option value={"distributer"}>Distributer</option>
            <option value={"retailer"}>Retailer</option>
          </Form.Select>
        </Form.Group>
        {!loading ? (
          <Button className="w-100" variant="primary" type="submit">
            Log In
          </Button>
        ) : (
          <Button className="w-100" variant="primary" type="submit" disabled>
            Logging In...
          </Button>
        )}
        <div className="d-grid justify-content-end">
          {/* <Button
            className="text-muted px-0"
            variant="link"
            onClick={handlePassword}
          >
            Forgot password?
          </Button> */}
          <Button
            className="text-muted px-0"
            variant="link"
            onClick={handleSignUpClick}
          >
            Don't have an account? Sign Up
          </Button>
        </div>
      </Form>
      {/* Footer */}
      {/* <div className="w-100 mb-2 position-absolute bottom-0 start-50 translate-middle-x text-white text-center">
        Made by Hendrik C | &copy;2022
      </div> */}
    </div>
  );
};

export default Login; 