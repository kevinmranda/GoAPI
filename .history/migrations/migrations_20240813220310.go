package migrations

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func SyncDatabase() {

}