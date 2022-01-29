package taskrepository

// Using this as the public facing 'repository'.
// Didn't do anything fancy, but if ever we wanted to replace boltdb we would only need to write the new implementation and change calls here.

func Update(taskName string) error {
	// Add task string to data store.
	return boltRepo.updateBolt(taskName)
}

func Delete(taskName string) error {
	// Remove task string from data store.
	return boltRepo.deleteBolt(taskName)
}

func Read() ([]string, error) {
	// Reads all task strings from data store.
	return boltRepo.readBolt()
}

func Clear() (err error) {
	return boltRepo.clearBolt()
}
