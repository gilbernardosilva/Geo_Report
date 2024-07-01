import React, { useState, useEffect } from "react";
import { Card, Row, Col, Button, Pagination, InputGroup, Form, Dropdown } from "react-bootstrap";
import { useAxiosWithToken } from "../../utils/api.js";
import CustomAuthorityNavbar from './../AuthorityNavbar';
import { useTranslation } from "react-i18next";
import cardIcon from '../../img/icons/cardIssues.svg'
import IssueDetailsModal from "../IssuesDetailModal/index.jsx";
import { useAuth } from "../../hooks/AuthContext.jsx";
import { formatDistanceToNow } from 'date-fns';
import { toast } from "react-toastify";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "react-toastify/dist/ReactToastify.css";
import "./index.css";
import NewEditIssueModal from "../NewEditIssueModal/index.jsx";


function AuthorityIssues() {
    const [issues, setIssues] = useState([]);
    const api = useAxiosWithToken();
    const [areas, setAreas] = useState([]);
    const [filterArea, setFilterArea] = useState(null);
    const [selectedIssue, setSelectedIssue] = useState(null);
    const [showModal, setShowModal] = useState(false);
    const { userInfo } = useAuth();
    const [totalPages, setTotalPages] = useState(1);
    const [currentPage, setCurrentPage] = useState(1);
    const [itemsPerPage, setItemsPerPage] = useState(10);
    const [startDate, setStartDate] = useState(null);
    const [endDate, setEndDate] = useState(null);

    const [showEditModal, setShowEditModal] = useState(false);


    const handleCloseEditModal = () => setShowEditModal(false);
    const handleEdit = (issue) => {
        setSelectedIssue(issue);
        setShowEditModal(true); 
    };

    const fetchIssues = async () => {
        try {
            const response = await api.get("admin/report", {
                params: {
                    page: currentPage,
                    limit: itemsPerPage,
                    start_date: startDate ? startDate.toISOString().split("T")[0] : "",
                    end_date: endDate ? endDate.toISOString().split("T")[0] : "",
                    area_id: filterArea ? filterArea : "",
                }
            });
            if (response.status === 200) {
                setIssues(response.data.reports);
                setTotalPages(response.data.total_pages);
            } else {
                toast.error("Error fetching issues");
            }
        } catch (err) {
            toast.error("Error fetching issues");
        }
    };


    useEffect(() => {
        const fetchAreas = async () => {
            try {
                const response = await api.get('admin/area');
                setAreas(response.data.areas);
            } catch (error) {
                toast.error("Failed to fetch areas.");
                console.error("Error fetching areas:", error);
            }
        }

        fetchAreas();

        if (userInfo) {
            fetchIssues();
        }
    }, [userInfo, api, currentPage, itemsPerPage, startDate, endDate, filterArea]);

    const handlePageChange = (page) => {
        setCurrentPage(page);
    };

    const handleViewDetails = (issue) => {
        setSelectedIssue(issue);
        setShowModal(true);
    };
    const handleAreaFilterChange = (areaId) => {
        setFilterArea(areaId);
    };

    const handleCloseModal = () => setShowModal(false);
    const { t } = useTranslation();

    return (
        <>
            <CustomAuthorityNavbar t={t}></CustomAuthorityNavbar>
            <h2 className="text-center mt-2" style={{ color: 'white', fontSize: '2rem' }}>My Issues</h2>
            <Row className="mb-3">
                <Col xs={12} md={8} className="d-flex align-items-center gap-2">

                    <InputGroup className="mb-3">
                        <InputGroup.Text className="inputGroup-issue">Start Date:</InputGroup.Text>
                        <DatePicker
                            selected={startDate}
                            onChange={(date) => setStartDate(date)}
                            selectsStart
                            startDate={startDate}
                            endDate={endDate}
                            dateFormat="yyyy-MM-dd"
                            className="form-control"
                        />
                        <InputGroup.Text className="inputGroup-issue">End Date:</InputGroup.Text>
                        <DatePicker
                            selected={endDate}
                            onChange={(date) => setEndDate(date)}
                            selectsEnd
                            startDate={startDate}
                            endDate={endDate}
                            minDate={startDate}
                            dateFormat="yyyy-MM-dd"
                            className="form-control"
                        />
                    </InputGroup>
                </Col>
                <Col xs={12} md={4} className="d-flex justify-content-end align-items-center">
                    <Dropdown onSelect={handleAreaFilterChange}>
                        <Dropdown.Toggle variant="secondary" id="area-filter-dropdown">
                            Filter by Area
                        </Dropdown.Toggle>
                        <Dropdown.Menu>
                            <Dropdown.Item eventKey={null}>All Areas</Dropdown.Item>
                            {areas.map((area) => (
                                <Dropdown.Item key={area.id} eventKey={area.id}>
                                    {area.name}
                                </Dropdown.Item>
                            ))}
                        </Dropdown.Menu>
                    </Dropdown>
                    <Form.Group>
                        <Form.Select className="pagination-container" value={itemsPerPage} onChange={(e) => setItemsPerPage(parseInt(e.target.value))}>
                            <option value={5}>5</option>
                            <option value={10}>10</option>
                            <option value={20}>20</option>
                        </Form.Select>
                    </Form.Group>
                </Col>
            </Row>
            <div className="d-flex flex-column align-items-center vh-100">

                <div className="d-flex justify-content-center flex-wrap gap-3">
                    {issues.map((issue) => (
                        <Card key={issue.id} className="dark-card" style={{ maxWidth: "80%" }}>
                            <Card.Header className="d-flex justify-content-between align-items-center">
                                <img src={cardIcon} alt="" className="img-fluid logo-icon" />
                                <Button disabled variant="info">
                                    {issue.report_status.status}
                                </Button>
                            </Card.Header>
                            <Card.Body>
                                <Row>
                                    <Col>
                                        <Card.Title>
                                            {t("name")}: {issue.name || "Not provided"}
                                        </Card.Title>
                                        <Card.Text>
                                            {t("type")}: {issue.report_type.name || "Not provided"}
                                        </Card.Text>
                                    </Col>

                                    <Col className="d-flex justify-content-end">
                                        <Button variant="primary" onClick={() => handleViewDetails(issue)}>
                                            {t("viewDetails")}
                                        </Button>
                                    </Col>
                                </Row>
                            </Card.Body>
                            <Card.Footer style={{ color: "grey" }} className="text-center">
                                <span>
                                    {formatDistanceToNow(new Date(issue.report_date))} {t("ago")}
                                </span>
                                <div>
                                    <Button variant="warning" className="me-2" onClick={() => handleEdit(issue)}>
                                        {t("edit")}
                                    </Button>
                                </div>
                            </Card.Footer>
                        </Card>

                    ))}

                    <Pagination className="mt-3">
                        {Array.from({ length: totalPages }, (_, index) => (
                            <Pagination.Item
                                key={index + 1}
                                active={index + 1 === currentPage}
                                onClick={() => handlePageChange(index + 1)}
                            >
                                {index + 1}
                            </Pagination.Item>
                        ))}
                    </Pagination>
                </div>
            </div>

            <IssueDetailsModal
                issue={selectedIssue}
                show={showModal}
                handleClose={handleCloseModal}
            />
            <NewEditIssueModal
                editMode={true}
                issue={selectedIssue}
                show={showEditModal}
                handleClose={handleCloseEditModal}
            />
        </>
    );
}

export default AuthorityIssues;
