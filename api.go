package request

// Get is request api with get method
func Get(config *Config) *Response {
	var res *Response
	if res = config.validate(); res != nil {
		return res
	}

	return config.Get()
}

// Post is request api with post method
func Post(config *Config) *Response {
	var res *Response
	if res = config.validate(); res != nil {
		return res
	}

	return config.Post()
}

// Delete is request api with delete method
func Delete(config *Config) *Response {
	if res := config.validate(); res != nil {
		return res
	}

	return config.Delete()
}

// Patch is request api with patch method
func Patch(config *Config) *Response {
	var res *Response

	if res = config.validate(); res != nil {
		return res
	}

	return config.Patch()
}

// Put is request api with put method
func Put(config *Config) *Response {
	var res *Response

	if res = config.validate(); res != nil {
		return res
	}

	return config.Patch()
}
