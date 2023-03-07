package data_service

import pipefiles "service/files_server/demo/pkg/pipefiles"

var (
	PipeService IPipeFilesServices = &pipeFilesHandlerServices{}
)

type pipeFilesHandlerServices struct{}

func (p *pipeFilesHandlerServices) PipeFilesSevice() {

	pipefiles.FileProcess()
}
