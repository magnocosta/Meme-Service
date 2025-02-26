package service

import (
  "meme_service/internal/app/meme/list/types"
	sharedtypes "meme_service/internal/shared/types"
)

var urls []string = []string{"a", "b", "c", "d"}

type fakeService struct {

}

func (r fakeService) List(input types.Input) (types.Outputs, error) {
  var outputs types.Outputs
  for _, v := range urls {
    item := sharedtypes.Meme{
      URL: v,
    }
    outputs = append(outputs, item)
  }
  return outputs, nil
}

func New() types.Service {
  return fakeService{}
}
