package repository
import(
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
type ParkingRecordRepositoryImpl interface{
	DB *grom.DB
}
func NewParkingRecordRepository(db *gorm.DB)*ParkingRecordRepositoryIpl{
	return &ParkingRecordRepository{DB:db}
}
func (r *ParkingRecordRepositoryImpl) Create(record *ParkingRecord) error {
	return r.DB.Create(record).Error
}

// FindByID 根据记录ID查找停车场记录
func (r *ParkingRecordRepositoryImpl) FindByID(id int) (*ParkingRecord, error) {
	var record ParkingRecord
	result := r.DB.First(&record, "record_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &record, nil
}

// Update 更新停车场记录
func (r *ParkingRecordRepositoryImpl) Update(record *ParkingRecord, id int) error {
	var existingRecord ParkingRecord
	result := r.DB.First(&existingRecord, "record_id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	updates := map[string]interface{}{}
	if record.VehicleID != existingRecord.VehicleID {
		updates["VehicleID"] = record.VehicleID
	}
	if record.SpaceID != existingRecord.SpaceID {
		updates["SpaceID"] = record.SpaceID
	}
	if record.LotID != existingRecord.LotID {
		updates["LotID"] = record.LotID
	}
	if record.StartTime != existingRecord.StartTime {
		updates["StartTime"] = record.StartTime
	}
	if record.EndTime != existingRecord.EndTime {
		updates["EndTime"] = record.EndTime
	}
	if record.Fee != existingRecord.Fee {
		updates["Fee"] = record.Fee
	}

	result = r.DB.Model(&existingRecord).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete 根据记录ID删除停车场记录
func (r *ParkingRecordRepositoryImpl) Delete(id int) error {
	return r.DB.Delete(&ParkingRecord{}, "record_id = ?", id).Error
}