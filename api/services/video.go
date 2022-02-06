package service

import (
	"learning-go/models"
	"learning-go/api/repositories"
)

type VideoService interface {
	Save(models.Video) models.Video
	Update(video models.Video)
	Delete(video models.Video)
	FindAll() []models.Video
}

type videoService struct {
	// videos []entity.Video
	videoRepository repository.VideoRepository
}

// Create a new instance for the video

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		// videos: []entity.Video{},
		videoRepository: repo,
	}
}

func (service *videoService) Save(video models.Video) models.Video {
	// service.videos = append(service.videos, video)
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) Update(video models.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video models.Video) {
	service.videoRepository.Delete(video)
}

func (service *videoService) FindAll() []models.Video {
	return service.videoRepository.FindAll()

}
