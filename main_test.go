package main

import (
	"testing"

	"github.com/coreos/go-semver/semver"
)

func stringToVersionSlice(stringSlice []string) []*semver.Version {
	versionSlice := make([]*semver.Version, len(stringSlice))
	for i, versionString := range stringSlice {
		versionSlice[i] = semver.New(versionString)
	}
	return versionSlice
}

func versionToStringSlice(versionSlice []*semver.Version) []string {
	stringSlice := make([]string, len(versionSlice))
	for i, version := range versionSlice {
		stringSlice[i] = version.String()
	}
	return stringSlice
}

func TestLatestVersions(t *testing.T) {
	testCases := []struct {
		versionSlice   []string
		expectedResult []string
		minVersion     *semver.Version
	}{
		{
			versionSlice:   []string{"1.8.11", "1.9.6", "1.10.1", "1.9.5", "1.8.10", "1.10.0", "1.7.14", "1.8.9", "1.9.5"},
			expectedResult: []string{"1.10.1", "1.9.6", "1.8.11"},
			minVersion:     semver.New("1.8.0"),
		},
		{
			versionSlice:   []string{"1.8.11", "1.9.6", "1.10.1", "1.9.5", "1.8.10", "1.10.0", "1.7.14", "1.8.9", "1.9.5"},
			expectedResult: []string{"1.10.1", "1.9.6"},
			minVersion:     semver.New("1.8.12"),
		},
		{
			versionSlice:   []string{"1.10.1", "1.9.5", "1.8.10", "1.10.0", "1.7.14", "1.8.9", "1.9.5"},
			expectedResult: []string{"1.10.1"},
			minVersion:     semver.New("1.10.0"),
		},
		{
			versionSlice:   []string{"2.2.1", "2.2.0"},
			expectedResult: []string{"2.2.1"},
			minVersion:     semver.New("2.2.1"),
		},
		// Implement more relevant test cases here, if you can think of any

		{
			versionSlice:   []string{"2.2.1", "2.2.0"},
			expectedResult: []string{""},
			minVersion:     semver.New("2.2.2"),
		},
		{
			versionSlice:   []string{"2.2.1", "2.1.0"},
			expectedResult: []string{""},
			minVersion:     semver.New("2.3.0"),
		},
		{
			versionSlice:   []string{"2.2.1", "2.2.0"},
			expectedResult: []string{""},
			minVersion:     semver.New("3.0.0"),
		},
		{
			versionSlice:   []string{"4.2.1", "3.5.7", "2.2.4", "1.2.3"},
			expectedResult: []string{"4.2.1", "3.5.7", "2.2.4"},
			minVersion:     semver.New("2.2.0"),
		},
		// test does not pass
		// {
		// 	versionSlice:   []string{"2.2.1", "2.2.0"},
		// 	expectedResult: []string{"2.2.1", "2.2.0"},
		// 	minVersion:     semver.New("2.1.0"),
		// },
	}

	test := func(versionData []string, expectedResult []string, minVersion *semver.Version) {
		stringSlice := versionToStringSlice(LatestVersions(stringToVersionSlice(versionData), minVersion))
		for i, versionString := range stringSlice {
			if versionString != expectedResult[i] {
				t.Errorf("Received %s, expected %s", stringSlice, expectedResult)
				return
			}
		}
	}

	for _, testValues := range testCases {
		test(testValues.versionSlice, testValues.expectedResult, testValues.minVersion)
	}
}
