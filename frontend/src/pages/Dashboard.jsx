import React, { useEffect, useState } from 'react'
import CoffeeTable from '../components/CoffeeTable';
import { getAllConffee, transferCoffee, updateAsset } from '../restService';
import { Button, Form, Modal } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';
import AllCoffeees from '../components/AllCoffeees';
import "../style/Dashboard.css";
import FarmerBackground from '../assests/farmerBackground.jpg';
import ProcessorBackground from '../assests/processorBackground.jpg'
import DistributorBackground from '../assests/distributorBackground.jpg'
import RetailerBackground from '../assests/retailerBackground.jpg'
import Spinner from 'react-bootstrap/Spinner';

const Dashboard = () => {
    const [user, setUser] = useState({})
    const [myCoffees, setMyCoffees] = useState([]);
    const [allCoffees, setAllCoffees] = useState([]);
    const [showMyCoffees, setShowMyCoffees] = useState(true);
    const [showAllCoffees, setShowAllCoffees] = useState(false);
    const [newOwner, setNewOwner] = useState("");
    const [currentCoffee, setCurrentCoffee] = useState({});
    const [openUpdateModal, setOpenUpdateModal] = useState(false);
    const [pageBackground, setPageBackground] = useState("");
    const [isLoading, setIsLoading] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        const ud = JSON.parse(localStorage.getItem("user"));
        if (ud.role == "farmer") {
            setNewOwner("Processor");
        } else if (ud.role == "processor") {
            setNewOwner("Distributer");
        } else if (ud.role == "distributer") {
            setNewOwner("Retailer");
        }
        setUser(ud);

        if (ud.role === "farmer") {
            setPageBackground(`${FarmerBackground}`);
        } else if (ud.role === "processor") {
            setPageBackground(`${ProcessorBackground}`);
        } else if (ud.role === "distributer") {
            setPageBackground(`${DistributorBackground}`);
        } else if (ud.role === "retailer") {
            setPageBackground(`${RetailerBackground}`);
        }

    }, [])

    const fetchData = () => {
        getAllConffee().then((resp) => {
            let data = resp.data.data
            console.log("ds - ", data);
            if (user.role == "farmer") {
                setAllCoffees(data.filter(c => c.producedBy == "farmer"));
            } else if (user.role == "processor") {
                setAllCoffees(data.filter(c => c.processedBy == "processor"));
            } else if (user.role == "distributer") {
                setAllCoffees(data.filter(c => c.distributedBy == "distributer"));
            }else if (user.role == "retailer") {
                setAllCoffees(data.filter(c => c.receivedBy == "retailer"));
            }
            data = data.filter(c => c.traceability.org == user.role)
            console.log(data);
            setMyCoffees(data);
        })
    }

    useEffect(() => {
        // getAllConffee().then((resp) => {
        //     let data = resp.data.data
        //     console.log("ds - ", data);
        //     if (user.role == "farmer") {
        //         setAllCoffees(data.filter(c => c.producedBy == "farmer"));
        //     } else if (user.role == "processor") {
        //         setAllCoffees(data.filter(c => c.processedBy == "processor"));
        //     } else if (user.role == "distributor") {
        //         setAllCoffees(data.filter(c => c.distributedBy == "distributor"));
        //     }
        //     data = data.filter(c => c.traceability.org == user.role)
        //     console.log(data);
        //     setMyCoffees(data);
        // })
        fetchData();
    }, [user])

    const onPressTransfer = (batchId) => {
        setIsLoading(true);
        const data = {
            batchId, currentowner: user.role, newowner: newOwner.toLowerCase()
        };
        console.log(data);
        transferCoffee(data).then((resp) => {
            console.log(resp);
            const data = myCoffees.filter(c => c.batchId != batchId);
            setIsLoading(false);
            setMyCoffees(data);
            fetchData();
        })
    }

    const onPressUpdate = (batchId) => {
        const c = myCoffees.filter(c => c.batchId == batchId);
        setCurrentCoffee(c[0]);
        setOpenUpdateModal(!openUpdateModal);
    }

    const makeUpdateApiCAll = (batchId, status, price) => {
        const data = {
            batchId, status, price
        }

        updateAsset(data).then(resp => {
            console.log("Update - ", resp);
            console.log(data);
            fetchData();
            setOpenUpdateModal(!openUpdateModal);
        })
    }

    const handleShowMyCoffees = () => {
        setShowMyCoffees(prevState => !prevState);
    }

    const handleShowAllCoffees = () => {
        setShowAllCoffees(prevState => !prevState);
    }

    return (
        <div className="mainContainer" style={{ backgroundImage: `linear-gradient(90deg, rgba(0,0,0,0.8) 0%, rgba(0,0,0,0.8) 100%), url(${pageBackground})` }}>
            <div className="container">
                <div className="row mt-5">
                    <div className="col">
                        <h1 className="text-danger">Welcome: {user.role}</h1>
                        <h2 className="text-success">{user.name}</h2>
                    </div>
                    <div className="col">
                        <div>
                            <button className='btn btn-primary mx-2' onClick={handleShowMyCoffees}>My Coffees</button>
                            <button className='btn btn-warning' onClick={handleShowAllCoffees}>All Coffees</button>
                        </div>
                        {/* {user.role == "farmer" && <button className="btn btn-primary mt-2" onClick={() => {
                        navigate("/farmer/createCoffee");
                    }}>
                        Create Coffee
                    </button>} */}
                    </div>
                </div>
                <div className="row mt-4">
                    {openUpdateModal && <UpdateModal user={user} coffee={currentCoffee} onClose={onPressUpdate} onPress={makeUpdateApiCAll} />}
                    {showMyCoffees && <CoffeeTable user={user} coffees={myCoffees} newOwner={newOwner} onPressTransfer={onPressTransfer} onPressUpdate={onPressUpdate} />}
                    {showAllCoffees && <AllCoffeees user={user} coffees={allCoffees} />}
                    {isLoading && <Loader />}
                </div>
            </div>
        </div>
    )
}

const UpdateModal = ({ coffee, onClose, onPress, user }) => {
    const [batchId, setBatchId] = useState("");
    const [status, setStatus] = useState("");
    const [price, setPrice] = useState("");

    useEffect(() => {
        console.log(coffee);
        setBatchId(coffee.batchId);
        if (user.role == "processor") {
            setStatus("PROCESSED");
        } else if (user.role == "distributer") {
            setStatus("PACKED");
        } else if (user.role == "retailer") {
            setStatus("READY TO CONSUME");
        }
        setPrice(coffee.price);
        console.log(status);
    }, [])

    const onChangeStatus = (event) => {
        console.log(event.target.value);
        setStatus(event.target.value);
    }

    return <div style={{
        display: 'block',
        width: 700,
        padding: 30, position: "absolute", zIndex: 99, backgroundColor: "white", left: "20%"
    }}
        className='shadow mx-auto'
    >        <Modal.Dialog>
            <Modal.Header closeButton onClick={onClose}>
                <Modal.Title>
                    Update Coffee
                </Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form className="shadow p-4 bg-white rounded" onSubmit={(event) => {
                    event.preventDefault();
                    onPress(batchId, status, price);
                }}>
                    <Form.Group className="mb-2" controlId="password">
                        <Form.Label>BatchId</Form.Label>
                        <Form.Control
                            type="text"
                            value={batchId}
                            placeholder="Password"
                            disabled
                        />
                    </Form.Group>
                    <Form.Group className="mb-2" controlId="password">
                        <Form.Label>Status</Form.Label>
                        <Form.Control
                            type="text"
                            value={status}
                            onChange={(e) => setPrice(e.target.value)}
                            disabled
                        />
                    </Form.Group>
                    <Form.Group className="mb-2" controlId="password">
                        <Form.Label>Price</Form.Label>
                        <Form.Control
                            type="number"
                            value={price}
                            onChange={(e) => setPrice(e.target.value)}
                        />
                    </Form.Group>
                    <Button className="w-25" variant="primary" type="submit">
                        Update
                    </Button>
                </Form>
            </Modal.Body>
        </Modal.Dialog>
    </div>
}

const Loader = () => {
    return (
        <div style={{ position: "absolute", backgroundColor: "black", opacity: 0.2, height: "80%", width: "100%", top: 0, left: 0, display: "flex", alignItems: "center", justifyContent: "center" }}>
            <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading...</span>
            </Spinner>
        </div>
    );
}


export default Dashboard;