import React from 'react'
import PropTypes from 'prop-types'

const AllCOffeees = ({ coffees, user }) => {
    return (
        <>
            <h3>All Coffees</h3>
            <table class="table table-striped table-hover">
                <thead>
                    <tr>
                        <th scope="col">Batch Id</th>
                        <th scope="col">Variety</th>
                        <th scope="col">Quantity</th>
                        <th scope="col">Origin</th>
                        <th scope="col">Status</th>
                        <th scope="col">Price</th>
                        <th scope="col">Produced By</th>
                        <th scope="col">Processed By</th>
                        <th scope="col">Distributed By</th>
                        <th scope="col">Retailed By</th>
                        <th scope="col">Current Owner</th>
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
                            <td>{coffee.producedBy}</td>
                            <td>{coffee.processedBy}</td>
                            <td>{coffee.distributedBy}</td>
                            <td>{coffee.receivedBy}</td>

                            <td>{coffee.traceability.org}</td>
                        </tr>)
                    }

                </tbody>
            </table>
        </>
    )
}

export default AllCOffeees