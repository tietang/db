package mysql

const (
	truncateArtist                             = "TRUNCATE TABLE `artist`"
	insertHayaoMiyazaki                        = "INSERT INTO `artist` (`name`) VALUES('Hayao Miyazaki')"
	insertIntoArtistWithPlaceholderReturningID = "INSERT INTO `artist` (`name`) VALUES(?)"
	selectFromArtistWhereName                  = "SELECT * FROM `artist` WHERE `name` = ?"
	updateArtistWhereName                      = "UPDATE `artist` SET `name` = ? WHERE `name` = ?"
	deleteArtistWhereName                      = "DELETE FROM `artis` WHERE `name` = $1"
)

//go:generate bash -c "sed s/ADAPTER/mysql/g ../internal/testing/adapter_benchmark.go > generated_benchmark_test.go"
