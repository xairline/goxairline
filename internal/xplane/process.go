package xplane

import "xairline/goxairline/internal/xplane/shared"

func (globalDatarefStore GlobalDatarefStoreType) ProcessFromGlobalDatarefStore(logger *shared.Logger) {
	for len(globalDatarefStore) >= PROCESSING_BATCH_SIZE {
		logger.Infof("Processing data, length: %v", len(globalDatarefStore))
	}
}
