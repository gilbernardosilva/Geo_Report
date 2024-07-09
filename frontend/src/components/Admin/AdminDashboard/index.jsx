import React, { useState, useEffect } from 'react';
import './index.css';
import { useAuth } from '../../../hooks/AuthContext';
import { useTranslation } from 'react-i18next';
import { useAxiosWithToken } from '../../../utils/api';
import ErrorHandler from '../../ErrorHandler/index';
import "react-toastify/dist/ReactToastify.css";
import CustomAdminNavbar from '../AdminNavbar/index';
import { Chart } from "react-google-charts"
import loading from "../../../img/loading.svg"
import luffy from "../../../img/noreports.png"
import { Card, CardBody, CardHeader, Col, Row } from 'react-bootstrap';

export const chartOptions = {
    titlePosition: 'none',
    is3D: true,
    backgroundColor: {
        fill: '#ffffff',
    },
    legend: { position: "bottom" },
    chartArea: {'width': '100%', 'height': '80%'},
  };


function AdminDashboard({ setToken }) {

    const { t } = useTranslation();
    const [error, setError] = useState(null);
    const { token } = useAuth();
    const api = useAxiosWithToken();
    const [reportTypeData, setReportTypeData] = useState([]);
    const [statusData, setStatusData] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            if (token !== null) {
                try {
                    const response = await api.get("admin");
                    console.log(response);
                    if (response.status === 200) {
                        console.log("has authorization");
                    } else {
                        setError(response);
                    }
                } catch (error) {
                    setError(error);
                }
            }
        };

        const fetchTypeData = async () => {
            if (token !== null) {
                try {
                    const response = await api.get("report/types/chart");
                    if (response.status === 200) {
                        const { labels, data } = response.data.data;
                        const chartData = labels.map((label, index) => [label, data[index]]);
                        setReportTypeData([
                            ["Type", "Count"],
                            ...chartData
                        ]);
                    } else {
                        setError(response);
                    }
                } catch (error) {
                    setError(error);
                }
            }
        };

        const fetchStatusData = async () => {
            try {
                const response = await api.get("report/status/chart");
                if (response.status === 200) {
                    const { labels, data } = response.data.data;
                    const chartData = labels.map((label, index) => [label, data[index]]);
                    setStatusData([
                        ["Status", "Count"],
                        ...chartData
                    ]);
                } else {
                    console.log(response);
                }
            } catch (error) {
                console.log(error);
            }
        };

        fetchData();
        fetchStatusData();
        fetchTypeData();
    }, [token]);


    const hasNoReports = (data) => {
        console.log(data)
        return data.length > 1 && data.slice(1).every(([label, count]) => count === 0);
    };

    return (
        <>
            {error ? (
                <ErrorHandler error={error} />
            ) : (
                <>
                    <CustomAdminNavbar setToken={setToken} t={t}></CustomAdminNavbar>

                    <div style={{padding: '30px'}}>
                        <Row className='d-flex'>
                            <Col className='d-flex text-center justify-content-center'>
                                <Card>
                                    <CardHeader>
                                        <h1>{t('typesIssues')}</h1>
                                    </CardHeader>
                                    <CardBody>
                                        {reportTypeData.length > 1 ? (
                                            hasNoReports(reportTypeData) ? (
                                                <div>
                                                    <p>{t('noReports')}</p>
                                                    <img src = {luffy} alt=':('/>
                                                </div>
                                            ) : (
                                            <Chart
                                                chartType="PieChart"
                                                data={reportTypeData}
                                                options={chartOptions}
                                                width={"100%"}
                                                height={"400px"}
                                            />
                                            )
                                        ) : (
                                            <img src={loading} alt="Loading..." />
                                        )}
                                    </CardBody>
                                </Card>
                            </Col>
                            <Col className='d-flex text-center justify-content-center'>
                                <Card>
                                    <CardHeader>
                                        <h1>{t('statusIssues')}</h1>
                                    </CardHeader>
                                    <CardBody>
                                        {statusData.length > 1 ? (
                                            hasNoReports(reportTypeData) ? (
                                                <div>
                                                    <p>{t('noReports')}</p>
                                                    <img src = {luffy} alt=':('/>
                                                </div>                                            
                                                ) : (
                                                <Chart
                                                    chartType="PieChart"
                                                    data={statusData}
                                                    options={chartOptions}
                                                    width={"100%"}
                                                    height={"400px"}
                                                />
                                            )
                                        ) : (
                                            <img src={loading} alt="Loading..." />
                                        )}
                                    </CardBody>
                                </Card>
                            </Col>  
                        </Row>
                    </div>
                </>
            )}
        </>
    )
}

export default AdminDashboard;