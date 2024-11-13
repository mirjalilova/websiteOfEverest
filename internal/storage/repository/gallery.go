package repository

import (
    "database/sql"
    "fmt"
    pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
    "strconv"
    "strings"
)

type GalleryRepository struct {
    db *sql.DB
}

func NewGalleryRepository(db *sql.DB) *GalleryRepository {
    return &GalleryRepository{db: db}
}

func (r *GalleryRepository) CreateGallery(req *pb.CreateGallery) (*pb.Void, error) {
    query := `INSERT INTO gallery (picture_url) VALUES ($1)`
    _, err := r.db.Exec(query, req.PictureUrl)
    if err != nil {
        return nil, err
    }
    return &pb.Void{}, nil
}

func (r *GalleryRepository) UpdateGallery(req *pb.UpdateGallery) (*pb.Void, error) {
    query := `UPDATE gallery SET`
    var args []interface{}
    var conditions []string

    if req.PictureUrl != "" {
        args = append(args, req.PictureUrl)
        conditions = append(conditions, "picture_url=$"+strconv.Itoa(len(args)))
    }

    conditions = append(conditions, "updated_at=NOW()")
    query += strings.Join(conditions, ", ")
    query += " WHERE id=$" + strconv.Itoa(len(args)+1) + " AND deleted_at=0"

    args = append(args, req.Id)

    _, err := r.db.Exec(query, args...)
    if err != nil {
        return nil, err
    }
    return &pb.Void{}, nil
}

func (r *GalleryRepository) DeleteGallery(req *pb.ById) (*pb.Void, error) {
    query := `UPDATE gallery SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1 AND deleted_at = 0`
    _, err := r.db.Exec(query, req.Id)
    if err != nil {
        return nil, err
    }
    return &pb.Void{}, nil
}

func (r *GalleryRepository) GetGalleryById(req *pb.ById) (*pb.GalleryRes, error) {
    res := &pb.GalleryRes{}
    query := `SELECT 
                id, 
                picture_url, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as created_at
            FROM 
                gallery 
            WHERE 
                id = $1 AND deleted_at = 0`
    
    err := r.db.QueryRow(query, req.Id).Scan(
        &res.Id,
        &res.PictureUrl,
        &res.CreatedAt,
    )
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("gallery item not found")
    } else if err != nil {
        return nil, err
    }
    return res, nil
}

func (r *GalleryRepository) GetGalleryList(req *pb.GetListGalleryReq) (*pb.GetListGalleryRes, error) {
    res := &pb.GetListGalleryRes{}
    query := `SELECT 
                COUNT(id) OVER() AS total_count, 
                id, 
                picture_url, 
                to_char(created_at, 'YYYY-MM-DD HH24:MI') as created_at 
            FROM 
                gallery 
            WHERE deleted_at = 0`
    
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var count int32
    for rows.Next() {
        var gallery pb.GalleryRes
        err := rows.Scan(&count, &gallery.Id, &gallery.PictureUrl, &gallery.CreatedAt)
        if err != nil {
            return nil, err
        }
        res.Galleries = append(res.Galleries, &gallery)
        res.TotalCount = count
    }

    if rows.Err() != nil {
        return nil, err
    }

    return res, nil
}
