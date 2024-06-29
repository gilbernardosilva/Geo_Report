import { Marker, Popup, useMapEvents } from "react-leaflet";

function LocationMarker({ position, setPosition, formData, setFormData }) {
  useMapEvents({
    click(e) {
      setPosition(e.latlng);
      setFormData((prevData) => ({
        ...prevData,
        latitude: e.latlng.lat,
        longitude: e.latlng.lng,
      }));
    },
  });

  return position === null || position === undefined ? null : ( 
    <Marker
      position={position}
      draggable={true}
      eventHandlers={{
        dragend(e) {
          const newPos = e.target.getLatLng();
          setPosition(newPos);
          setFormData((prevData) => ({
            ...prevData,
            latitude: newPos.lat,
            longitude: newPos.lng,
          }));
        },
      }}
    >
      <Popup>
        Latitude: {position.lat} <br /> Longitude: {position.lng}
      </Popup>
    </Marker>
  );
}

export default LocationMarker;

