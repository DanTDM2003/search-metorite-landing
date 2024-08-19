package locator

import "sync"

type ServiceLocator struct {
	services map[string]interface{}
	mu       sync.RWMutex
}

var instance *ServiceLocator
var once sync.Once

func GetServiceLocator() *ServiceLocator {
	once.Do(func() {
		instance = &ServiceLocator{
			services: make(map[string]interface{}),
		}
	})
	return instance
}

func (s *ServiceLocator) RegisterService(name string, service interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.services[name] = service
}

func (s *ServiceLocator) GetService(name string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.services[name]
}
