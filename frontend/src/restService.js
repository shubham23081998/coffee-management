import axios from "axios";
import { getAllConffeeResp, loginInResp } from "./data";

const baseUrl = "http://localhost:8002";

export const signUp = async (data) => {
    console.log(data);
    return await axios.post(baseUrl + "/signup", data)
}

export const login = async (email, password, role) => {
    const data = {
        email, password, role
    }
    console.log(data);
    return await axios.post(baseUrl + "/login", data)
}

export const getAllConffee = async () => {
    return await axios.get(baseUrl + "/getAllCoffee");
}

export const transferCoffee = async (data) => {
    console.log(data);
    return await axios.post(baseUrl + "/trasnsfercoffee", data);
}

export const createCoffee = async (data) => {
    console.log(data);
    return await axios.post(baseUrl + "/createCoffee", data)
}

export const updateAsset = async (data) => {
    console.log(data);
    return await axios.post(baseUrl + "/updatecoffee", data);
}