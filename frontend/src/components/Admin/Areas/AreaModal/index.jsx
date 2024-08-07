import React, { useState, useEffect } from "react";
import { Modal, Button, Form } from "react-bootstrap";
import { MapContainer, TileLayer, Marker, Popup, useMapEvents } from "react-leaflet";
import L, { map } from "leaflet";
import "leaflet/dist/leaflet.css";
import { toast } from "react-toastify";
import { useAxiosWithToken } from "../../../../utils/api";
import LocationMarker from "../../../LocationMarker";


function AreaModal({ show, onHide, editMode, area, onSubmit, fetchAreas }) {
  const api = useAxiosWithToken();
  const [formData, setFormData] = useState({
    name: "",
    latitude: 0,
    longitude: 0,
    radius: 0,
  });
  const [mapPosition, setMapPosition] = useState({
    lat: area?.latitude || 0,
    lng: area?.longitude || 0,
  });



  useEffect(() => {
    if (show) {
      if (area) {
        setFormData(area);
        setMapPosition({ lat: area.latitude, lng: area.longitude });
      } else {
        setFormData({
          name: "",
          latitude: 0,
          longitude: 0,
          radius: 0,
        });
        setMapPosition(null);
      }
    }
  }, [show, area]);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: name === "radius" ? parseFloat(value) : value,
    }));
  };

  const handleSubmit = async (event) => {
    debugger;
    event.preventDefault();
    const dataToSend = { ...formData, latitude: mapPosition.lat, longitude: mapPosition.lng };
    try {
      if (editMode) {
        const response = await api.put('admin/area/' + area.id, dataToSend);
        if (response.status === 200) {
          toast.success("Area updated successfully");
          fetchAreas();
          onHide();
        } else {
          toast.error("Failed to updated area");
          console.error("Error updating area:", response.data);
        }
      } else {
        const response = await api.post('admin/area', dataToSend);
        if (response.status === 200) {
          toast.success("Area created successfully");
          onHide();
          fetchAreas();
        } else {
          toast.error("Failed to create area");
          console.error("Error creating area:", response.data);
        }
      }
    } catch (error) {
      toast.error("Failed to save area");
      console.error("Error saving area:", error);
    }
  };
  return (
    <Modal show={show} onHide={onHide}>
      <Modal.Header closeButton>
        <Modal.Title>{editMode ? "Edit Area" : "Create Area"}</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form onSubmit={handleSubmit}>
          <Form.Group controlId="name">
            <Form.Label>Name</Form.Label>
            <Form.Control
              type="text"
              name="name"
              value={formData.name}
              onChange={handleChange}
              required
            />
          </Form.Group>
          <Form.Group controlId="radius">
            <Form.Label>Radius (meters)</Form.Label>
            <Form.Control
              type="number"
              name="radius"
              value={formData.radius}
              onChange={handleChange}
              required
            />
          </Form.Group>
          <MapContainer
            center={[area.latitude, area.longitude]}
            zoom={4}
            scrollWheelZoom={true}
            style={{ height: "300px", width: "100%" }}

          >
            <Marker
              position={[mapPosition.lat, mapPosition.lng]}
            >
            </Marker>
            <TileLayer
              attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
              url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            <LocationMarker setFormData={setFormData} formData={formData} setPosition={setMapPosition} />
          </MapContainer>

          

          <Button variant="primary" type="submit" className="mt-3">
            {editMode ? "Save Changes" : "Create Area"}
          </Button>
        </Form>
      </Modal.Body>
    </Modal>
  );
}

export default AreaModal;
