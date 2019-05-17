package jenkins

// Job -- jenkins job
type Job struct {
	Name      string
	ConfigXML string
}

// Delete -- 删除job
func Delete() error {
	return nil
}

// Create -- 创建job
func Create() error {
	return nil
}
