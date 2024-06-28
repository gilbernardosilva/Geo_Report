import React, { useState, useEffect } from 'react';
import { Button, Modal, Form, Table, Container } from 'react-bootstrap';
import { useAxiosWithToken } from '../../../utils/api';
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import AreaModal from './AreaModal';
import CustomAdminNavbar from '../AdminNavbar';
import { useTranslation } from 'react-i18next';
function AdminArea() {
    const {t} = useTranslation();
    const api = useAxiosWithToken();
    const [areas, setAreas] = useState([]);
    const [showModal, setShowModal] = useState(false);
    const [editMode, setEditMode] = useState(false);
    const [formData, setFormData] = useState({
        name: "",
        latitude: 0,
        longitude: 0,
        radius: 0,
    });
    const [selectedArea, setSelectedArea] = useState({
        id: null,
        name: '',
        latitude: 0,
        longitude: 0,
        radius: 0,
    });

    useEffect(() => {
        fetchAreas();
    }, []);

  async function fetchAreas() {
    try {
      const response = await api.get('admin/area');
      setAreas(response.data.areas);
    } catch (error) {
      toast.error("Failed to fetch areas.");
      console.error("Error fetching areas:", error);
    }
  }

 const handleClose = () => {
        setFormData({
            name: "",
            latitude: 0,
            longitude: 0,
            radius: 0,
        });
        setShowModal(false);
        setEditMode(false);
        setSelectedArea(null); 
    };

  const handleEdit = (area) => {
    setSelectedArea(area);
    setEditMode(true);
    setShowModal(true);
  };

  const handleDelete = async (areaId) => {
    try {
      await api.delete(`admin/area/${areaId}`);
      fetchAreas();
      toast.success("Area deleted successfully");
    } catch (error) {
      toast.error("Failed to delete area");
      console.error("Error deleting area:", error);
    }
  };

  const handleCreate = () => {
    setSelectedArea(null);
    setEditMode(false);
    setShowModal(true);
  };

  const handleSubmit = async (event) => {
    debugger;
    event.preventDefault();
    try {
      if (editMode) {
        await api.put(`admin/area/${selectedArea.id}`, formData);
        toast.success("Area updated successfully");
      } else {
        await api.post('admin/area', formData);
        toast.success("Area created successfully");
      }
      fetchAreas();
      setShowModal(false);
    } catch (error) {
      toast.error("Failed to save area");
      console.error("Error saving area:", error);
    }
  };

  return (
    <>
    <CustomAdminNavbar t={t}></CustomAdminNavbar>
    <Container>
      <h2>Area Management</h2>
      <Button onClick={handleCreate}>Create New Area</Button>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Latitude</th>
            <th>Longitude</th>
            <th>Radius</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {areas.map((area) => (
            <tr key={area.id}>
              <td>{area.id}</td>
              <td>{area.name}</td>
              <td>{area.latitude}</td>
              <td>{area.longitude}</td>
              <td>{area.radius}</td>
              <td>
                <Button variant="warning" onClick={() => handleEdit(area)}>Edit</Button>
                <Button variant="danger" onClick={() => handleDelete(area.id)}>Delete</Button>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>

      {/* Modal Component */}
      <AreaModal
        show={showModal}
        onHide={handleClose}
        editMode={editMode}
        area={selectedArea}
        onSubmit={handleSubmit}
        fetchAreas={fetchAreas}
      />
    </Container>
    </>
  );
}

export default AdminArea;