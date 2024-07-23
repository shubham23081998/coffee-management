import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom';
import { createCoffee } from '../restService';

const CreateCoffee = () => {
    const [batchId, setBatchId] = useState("");
    const [variety, setVariety] = useState("");
    const [quantity, setQuantity] = useState("");
    const [origin, setOrigin] = useState("");
    const [price, setPrice] = useState("");
    const [createCoffeeLoading, setCreateCoffeeLoading] = useState(false);
    const navigate = useNavigate();
    const onSubmit = (event) => {
        setCreateCoffeeLoading(true);
        event.preventDefault();
        const data = {
            batchId: batchId,
            variety: variety,
            quantity: quantity,
            origin: origin,
            price: price,
            producedBy: JSON.parse(localStorage.getItem("user")).role
        }

        createCoffee(data).then(resp => {
            console.log("REsp - ", resp);
            navigate("/dashboard");
            setCreateCoffeeLoading(false);
        })
        console.log("Data - ", data);
    }
    return (
        <div className="container">
            <form className='w-25 mx-auto mt-4'>
                <div class="form-group">
                    <label for="batchId">Batch ID</label>
                    <input type="number" class="form-control" placeholder="Enter Batch ID" onChange={(e) => setBatchId(e.target.value)} />
                </div>
                <div class="form-group mt-1">
                    <label for="variety">Variety</label>
                    <input type="text" class="form-control" placeholder="Enter variety" onChange={(e) => setVariety(e.target.value)} />
                </div>
                <div class="form-group mt-1">
                    <label for="quantity">Quantity</label>
                    <input type="number" class="form-control" placeholder="Enter quantity" onChange={(e) => setQuantity(e.target.value)} />
                </div>
                <div class="form-group mt-1">
                    <label for="origin">Origin</label>
                    <input type="text" class="form-control" placeholder="Enter origin" onChange={(e) => setOrigin(e.target.value)} />
                </div>
                <div class="form-group mt-1">
                    <label for="quantity">Price</label>
                    <input type="number" class="form-control" placeholder="Enter price" onChange={(e) => setPrice(e.target.value)} />
                </div>
                {/* <button type="submit" class="btn btn-primary mt-3" onClick={onSubmit}>Create</button> */}
                {!createCoffeeLoading ? (
                    <button type="submit" class="btn btn-primary mt-3" onClick={onSubmit}>Create</button>
                ) : (
                    <button type="submit" class="btn btn-primary mt-3" onClick={onSubmit} disabled>Creating...</button>
                )}
            </form>
        </div>

    )
}

export default CreateCoffee