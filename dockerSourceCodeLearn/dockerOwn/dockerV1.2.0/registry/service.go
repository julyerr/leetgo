package registry

import (
	"github.com/docker/docker/engine"
)

// 'auth' : Authenticate against the public registry
// 'search' : Search for images on the public registry
// 'pull' : Download images from any registry (TODO)
// 'push' : Upload images to any registry (TODO)
type Service struct{

}

func (s *Service) Auth(job *engine.Job) engine.Status{
	var (
		err error
		authConfig = &AuthConfig{}
		)
	job.GetenvJson("authConfig",authConfig)
	if addr := authConfig.ServerAddress;addr != "" && addr != IndexServerAddress(){
		addr,err = ExpandAndVerifyRegistryUrl(addr)
		if err != nil{
			return job.Error(err)
		}
		authConfig.ServerAddress = addr
		status,err := Login(authConfig,HTTPRequestFactory(nil))
		if err != nil{
			return job.Error(err)
		}
		job.Printf("%s\n",status)
		return engine.StatusOK
	}
}

