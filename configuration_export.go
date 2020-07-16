package zabbix

type ConfiguraOption struct {
	// ids of host groups to export
	Groups []string `json:"groups,omitempty"`
	// ids of hosts to export
	Hosts []string `json:"hosts,omitempty"`
	// ids of images to export
	Images []string `json:"images,omitempty"`
	// ids of maps to export
	Maps []string `json:"maps,omitempty"`
	// ids of screens to export
	Screens []string `json:"screens,omitempty"`
	// ids of templaates to export
	Templates []string `json:"templates,omitempty"`
	// ids of value maps to export
	ValueMaps []string `json:"valueMaps,omitempty"`
}

type ConfigurationParamsRequest struct {
	// base options
	GetParameters
	// format in which the data must be exported, json =JSON, xml =XML
	Format string `json:"format"`
	// objects to be exported
	Options ConfiguraOption `json:"options"`
}

type ConfigurationExportResponse string

func (c *Session) ConfiguraExport(params ConfigurationParamsRequest) (ConfigurationExportResponse, error) {
	var respData ConfigurationExportResponse
	return respData, c.Get("configuration.export", params, &respData)
}
