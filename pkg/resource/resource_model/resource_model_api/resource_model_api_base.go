package resource_model_api

func (resourceModelApi *ResourceModelApi) RecreateDatabase() error {
	if err := resourceModelApi.DropDatabase(); err != nil {
		return err
	}

	if err := resourceModelApi.CreateDatabase(); err != nil {
		return err
	}

	return nil
}

func (resourceModelApi *ResourceModelApi) DropDatabase() error {
	return nil
}

func (resourceModelApi *ResourceModelApi) CreateDatabase() error {
	return nil
}
