package backgroundjob

import "system_service/types"

type BackgroundJobModule struct {
	ipc *types.InternalProviderControl
}

func NewBackgroundJobModule(ipc *types.InternalProviderControl) *BackgroundJobModule {
	return &BackgroundJobModule{ipc: ipc}
}

func (m *BackgroundJobModule) StartJobsAll() {
	// Запуск синхронизации базы данных
	m.StartDatabaseSync()
}
