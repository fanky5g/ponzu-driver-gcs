package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UtilTestSuite struct {
	suite.Suite
}

func (s *UtilTestSuite) TestParsePathWithGSPathSyntax() {
	bucket, key, err := parsePath("gs://bucket-34ff28a4-b823-4c05-8367-8d760927938b/test_file.pdf")
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), bucket, "bucket-34ff28a4-b823-4c05-8367-8d760927938b")
		assert.Equal(s.T(), key, "test_file.pdf")
	}
}

func (s *UtilTestSuite) TestParsePathWithBucketAndKey() {
	bucket, key, err := parsePath("bucket-34ff28a4-b823-4c05-8367-8d760927938b/test_file.pdf")
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), bucket, "bucket-34ff28a4-b823-4c05-8367-8d760927938b")
		assert.Equal(s.T(), key, "test_file.pdf")
	}
}

func (s *UtilTestSuite) TestParsePathWithNestedBucketAndKey() {
	bucket, key, err := parsePath("bucket-34ff28a4-b823-4c05-8367-8d760927938b/certificates/2024/test_file.pdf")
	if assert.NoError(s.T(), err) {
		assert.Equal(s.T(), bucket, "bucket-34ff28a4-b823-4c05-8367-8d760927938b")
		assert.Equal(s.T(), key, "certificates/2024/test_file.pdf")
	}
}

func (s *UtilTestSuite) TestParsePathThrowsErrorForInvalidPaths() {
	_, _, err := parsePath("test_file.pdf")
	assert.EqualError(s.T(), err, ErrInvalidPath.Error())
}

func TestUtils(t *testing.T) {
	suite.Run(t, new(UtilTestSuite))
}
