import React, { useState, useEffect } from "react";
import { Form, Button, Table, Container } from 'react-bootstrap';
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useTranslation } from "react-i18next";
import { useAxiosWithToken } from "../../../utils/api.js";
import CustomAdminNavbar from "../AdminNavbar/index.jsx";


function ReportTypePage() {
    const [newReportTypeName, setNewReportTypeName] = useState("");
    const { t } = useTranslation();
    const api = useAxiosWithToken();
    const [reportTypes, setReportTypes] = useState([]);

    useEffect(() => {
        const fetchReportTypes = async () => {
            try {
                const response = await api.get("/reportTypes");
                setReportTypes(response.data.report_types);
            } catch (error) {
                console.error("Error fetching report types:", error);
                toast.error("Failed to fetch report types");
            }
        };
        fetchReportTypes();
    }, []);


    const handleCreate = async () => {
        if (newReportTypeName.trim() !== "") {
            try {
                const response = await api.post("admin/report/reportType", { name: newReportTypeName });
                if(response.status === 200){
                setReportTypes([...reportTypes, response.data.reporty_type]);
                setNewReportTypeName(""); // Clear the input field
                toast.success("New report type created successfully!");
                } else {
                    toast.error("Failed to create report type");
                }
            } catch (error) {
                console.error("Error creating report type:", error);
                toast.error("Failed to create report type");
            }
        }
    };

    return (
        <>
            <CustomAdminNavbar t={t}></CustomAdminNavbar>
            <Container>
                <h2>{t('types')}</h2>
                <Form.Group controlId="newReportTypeName">
                    <Form.Control
                        type="text"
                        placeholder="Enter new report type name"
                        value={newReportTypeName}
                        onChange={(e) => setNewReportTypeName(e.target.value)}
                    />
                </Form.Group>
                <Button className="mb-3 mt-3" onClick={handleCreate}>Create Type</Button>

                <Table striped bordered hover>
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Name</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {reportTypes.map((rt) => (
                            <tr key={rt.id}>
                                <th>{rt.id}</th>
                                <th>{rt.name}</th>
                                <th>
                                    <Button disabled>{t('delete')}</Button>
                                </th>
                            </tr>
                        ))}
                    </tbody>
                </Table>
            </Container>
        </>
    );
}

export default ReportTypePage;
