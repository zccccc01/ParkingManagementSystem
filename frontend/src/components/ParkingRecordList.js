// src/components/ParkingRecordList.js
import React from 'react';

const ParkingRecordList = ({ records }) => {
  return (
    <ul className="parking-record-list">
      {records.map((record, index) => (
        <li key={index}>
          <p>Parking Lot: {record.parkingLot}</p>
          <p>Duration: {record.duration}</p>
          <p>Cost: ${record.cost}</p>
        </li>
      ))}
    </ul>
  );
};

export default ParkingRecordList;