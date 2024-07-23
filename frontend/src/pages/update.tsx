import React from 'react'
import PropTypes from 'prop-types'

const update = props => {
    return (
        <div>
            <div className="mb-3">
                <label htmlFor="exampleFormControlInput1" className="form-label">Email address</label>
                <input type="email" className="form-control" id="exampleFormControlInput1" placeholder="name@example.com" />
            </div>
            <div className="mb-3">
                <label htmlFor="exampleFormControlTextarea1" className="form-label">Example textarea</label>
            </div>
        </div>
    )
}

update.propTypes = {}

export default update