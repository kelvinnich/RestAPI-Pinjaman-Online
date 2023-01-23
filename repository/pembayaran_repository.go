package repository

import (
	"pinjaman-online/model"
	"time"

	"gorm.io/gorm"
)

type PembayaranRepository interface {
	CreatePembayaranRepository(pembayaran *model.Pembayaran) error
	FindByIdRepository(id int) (*model.Pembayaran, error)
	UpdatePembayaranRepository(id int, pembayaran *model.Pembayaran) error
	DeletePembayaranRepository(id int) error
	ListPembayaranRepository() ([]*model.Pembayaran, error)
	GetPembayaranPerBulanRepository(pinjamanID int) (int, error)
	GetTotalPembayaranRepository(pinjamanID int) (int, error)
	GetJatuhTempoPembayaranRepository(pinjamanID int) ([]time.Time, error)
}

type pembayaranConnection struct {
	db *gorm.DB
}


func NewPembayaranRepository(db *gorm.DB)PembayaranRepository{
	return &pembayaranConnection{
		db: db,
	}
}

func(db *pembayaranConnection)CreatePembayaranRepository(pembayaran *model.Pembayaran) error{
	if err := db.db.Create(pembayaran).Error; err != nil {
		return err
	}

	if err := db.init(); err != nil{
		return err
	}

	return nil
}

func(db *pembayaranConnection )UpdatePembayaranRepository(id int, pembayaran *model.Pembayaran) error{
	if err := db.db.Model(&model.Pembayaran{}).Where("id = $1", id).Updates(pembayaran).Error; err != nil {
		return err
	}

	if err := db.init(); err != nil{
		return err
	}

	return nil
}

func(db *pembayaranConnection)FindByIdRepository(id int) (*model.Pembayaran, error){
	var pembayaran model.Pembayaran
	if err := db.db.First(&pembayaran, id).Error; err != nil {
		return nil,err
	}

	return &pembayaran,nil
}

func(db *pembayaranConnection)DeletePembayaranRepository(id int) error{
	if err := db.db.Where("id = $1", id).Delete(&model.Pembayaran{}).Error; err != nil {
		return err
	}
	return nil
}

func(db *pembayaranConnection)ListPembayaranRepository() ([]*model.Pembayaran, error){
    var pembayarans []*model.Pembayaran
    if err := db.db.Find(&pembayarans).Error; err != nil {
        return nil, err
    }
    return pembayarans, nil
}

func (db *pembayaranConnection) GetPembayaranPerBulanRepository(pinjamanID int) (int, error) {
	var pinjaman model.Pinjaman
	if err := db.db.First(&pinjaman, pinjamanID).Error; err != nil {
			return 0, err
	}
	jumlahPinjaman := pinjaman.Jumlah
	sukuBunga := pinjaman.SukuBunga
	durasiPinjaman := pinjaman.Durasi
	pembayaranPerBulan := (jumlahPinjaman * sukuBunga) / (12 * 100) + (jumlahPinjaman / durasiPinjaman)
	return pembayaranPerBulan, nil
}

func (db *pembayaranConnection) GetTotalPembayaranRepository(pinjamanID int) (int, error) {
	var pinjaman model.Pinjaman
	if err := db.db.First(&pinjaman, pinjamanID).Error; err != nil {
	return 0, err
	}
	jumlahPinjaman := pinjaman.Jumlah
	sukuBunga := pinjaman.SukuBunga
	durasiPinjaman := pinjaman.Durasi
	pembayaranPerBulan := (jumlahPinjaman * sukuBunga) / (12 * 100) + (jumlahPinjaman / durasiPinjaman)
	totalPembayaran := pembayaranPerBulan * durasiPinjaman
	return totalPembayaran, nil
	}

func (db *pembayaranConnection) GetJatuhTempoPembayaranRepository(pinjamanID int) ([]time.Time, error) {
    var pembayarans []*model.Pembayaran
    var jatuhTempo []time.Time
    var pinjaman model.Pinjaman

    if err := db.db.First(&pinjaman, pinjamanID).Error; err != nil {
        return nil, err
    }

    if err := db.db.Where("pinjaman_id = $1", pinjamanID).Find(&pembayarans).Error; err != nil {
        return nil, err
    }

    for _, pembayaran := range pembayarans {
        jatuhTempo = append(jatuhTempo, pembayaran.Tanggal_Pembayaran.AddDate(0, int(pinjaman.Durasi), 0))
    }

    return jatuhTempo, nil
}

func(db *pembayaranConnection) init() error {

	triggerSQL := `
	CREATE TRIGGER update_payment_status
	AFTER INSERT ON pembayarans
	FOR EACH ROW
	BEGIN
	UPDATE pembayarans
	SET status_pembayaran = true
	WHERE id = NEW.id;
	END;
	`

	if err := db.db.Exec(triggerSQL).Error; err != nil {
			return err
	}
	return nil
}





