package service

import (
	"learning-go/entity"
	"learning-go/repositories"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
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

func (service *videoService) Save(video entity.Video) entity.Video {
	// service.videos = append(service.videos, video)
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()

}
