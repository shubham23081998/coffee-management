import React, { useState } from 'react'
import PropTypes from 'prop-types'
import { useNavigate } from 'react-router-dom';

const CoffeeTable = ({ coffees, newOwner, onPressTransfer, onPressUpdate, user }) => {
    const [transferLoading, setTransferLoading] = useState(false);
    const navigate = useNavigate();
    return (
        <>
            <div style={{ display: "flex", justifyContent: "space-between", alignContent: "center", margin: "15px 0" }}>
                <h2>My Coffees</h2>
                {user.role == "farmer" && <button className="btn btn-success" onClick={() => {
                    navigate("/farmer/createCoffee");
                }}>
                    Create Coffee
                </button>}
            </div>
            <table class="table table-striped table-hover">
                <thead>
                    <tr>
                        <th scope="col">Batch Id</th>
                        <th scope="col">Variety</th>
                        <th scope="col">Quantity</th>
                        <th scope="col">Origin</th>
                        <th scope="col">Status</th>
                        <th scope="col">Price</th>
                        <th scope="col">Current Owner</th>
                        <th scope="col">{user.role !== "retailer" ? "Actions" : "Status"}</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        coffees.map((coffee, index) => <tr>
                            <th scope="row">{coffee.batchId}</th>
                            <td>{coffee.variety}</td>
                            <td>{coffee.quantity}</td>
                            <td>{coffee.origin}</td>
                            <td>{coffee.status}</td>
                            <td>{coffee.price}</td>
                            <td>{coffee.traceability.org}</td>
                            <td>
                                {user.role !== "retailer"
                                    ?
                                    <div>
                                        {user.role != "farmer" && <button className='btn btn-primary m-2 my-0' onClick={() => onPressUpdate(coffee.batchId)}>Update</button>}
                                        {/* <button className='btn btn-success my-0' onClick={() => { setTransferLoading(true); onPressTransfer(coffee.batchId) }}>Transfer {`To ${newOwner}`}</button> */}

                                        <button className='btn btn-success my-0' onClick={() => { setTransferLoading(true); onPressTransfer(coffee.batchId) }}>Transfer {`To ${newOwner}`}</button>
                                    </div>
                                    :
                                    <div>
                                         {user.role != "farmer" && <button className='btn btn-primary m-2 my-0' onClick={() => onPressUpdate(coffee.batchId)}>Update</button>}
                                        <button  className='btn btn-success my-0' disabled>Ready To Sell</button></div>
                                }
                            </td>
                        </tr>)
                    }

                </tbody>
            </table>
        </>)
}

CoffeeTable.propTypes = {}

export default CoffeeTable