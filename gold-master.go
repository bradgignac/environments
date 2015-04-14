stats.Time("myservice.operation.v1", func() {
	operation.VersionOne()
})

stats.Time("myservice.operation.v2", func() {
	operation.VersionTwo()
})
