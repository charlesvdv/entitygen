package entitygen

import (
	"errors"
	"fmt"
	"io"
	"os"

	gomigratesource "github.com/golang-migrate/migrate/v4/source"
)

var ErrEndOfIteration = errors.New("end of iterations")

// MigrationSource abstracts the location and the handling of migration definitions.
type MigrationSource interface {
	// Returns the next migration script in the correct order. If the end of the scripts is reached,
	// this method returns `ErrEndOfIteration`.
	Next() (io.ReadCloser, error)

	// Close the underlying source if needed.
	Close() error
}

// The driver needs to be opened before it is passed here.
func NewGoMigrateSource(driver gomigratesource.Driver) (GoMigrateSource, error) {
	return GoMigrateSource{
		driver:  driver,
		version: nil,
	}, nil
}

type GoMigrateSource struct {
	driver  gomigratesource.Driver
	version *uint
}

func (s *GoMigrateSource) Next() (io.ReadCloser, error) {
	err := s.moveToNextVersion()
	if err != nil {
		return nil, err
	}

	reader, _, err := s.driver.ReadUp(*s.version)
	if errors.Is(err, os.ErrNotExist) {
		// Weird but we may have some migration steps where only downgrades are defined.
		return s.Next()
	}
	if err != nil {
		return nil, fmt.Errorf("go migrate source: failed to read version '%d': %w", *s.version, err)
	}

	return reader, nil
}

func (s *GoMigrateSource) moveToNextVersion() error {
	var version uint
	var err error

	if s.version == nil {
		version, err = s.driver.First()
		if err != nil {
			return fmt.Errorf("go migrate source: failed to init driver: %w", err)
		}
	} else {
		version, err = s.driver.Next(*s.version)
		if errors.Is(err, os.ErrNotExist) {
			return ErrEndOfIteration
		} else if err != nil {
			return fmt.Errorf("go migrate source: failed to advance iterator: %w", err)
		}
	}

	s.version = &version
	return nil
}

func (s *GoMigrateSource) Close() error {
	return s.driver.Close()
}
