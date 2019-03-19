package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
)

// type Version struct {
// 	Major      int64
// 	Minor      int64
// 	Patch      int64
// 	PreRelease PreRelease
// 	Metadata   string
// }

// LatestVersions returns a sorted slice with the highest version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	if releases == nil {
		fmt.Println("No Releases found")
	}
	if minVersion == nil {
		fmt.Println("No minVersion found")
	}
	// This is just an example structure of the code, if you implement this interface, the test cases in main_test.go are very easy to run
	var versionSlice []*semver.Version

	// sort the releases slice in Version AscendingOrder
	sort.Sort(AscendingOrder(releases))

	// Check if
	for _, r := range releases {
		if r.Compare(*minVersion) >= 0 {
			if versionSlice == nil {
				//Append versionSlice with r when versionSlice is NIL
				versionSlice = append(versionSlice, r)
			} else {
				sliceR := r.Slice()
				sliceV := versionSlice[len(versionSlice)-1].Slice()
				// append versionSlice with r when
				if sliceR[1] != sliceV[1] {
					versionSlice = append(versionSlice, r)
				}
			}
		} else {
			// Rest of the Release slice is smaller than minVersion, break from loop
			break
		}
	}
	sort.Sort(AscendingOrder(versionSlice))

	return versionSlice
}

// AscendingOrder sorter
type AscendingOrder []*semver.Version

// return length of version
func (a AscendingOrder) Len() int { return len(a) }

// boolean function that returns true if a[i] < a[j]
func (a AscendingOrder) Less(i, j int) bool { return a[i].LessThan(*a[j]) }

// swaps the placement of a[i] with a[j]
func (a AscendingOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Sort sorts the given slice of Version
func Sort(ascendingOrder []*semver.Version) { sort.Sort(AscendingOrder(ascendingOrder)) }

// Here we implement the basics of communicating with github through the library as well as printing the version
// You will need to implement LatestVersions function as well as make this application support the file format outlined in the README
// Please use the format defined by the fmt.Printf line at the bottom, as we will define a passing coding challenge as one that outputs
// the correct information, including this line
func main() {
	// Github
	client := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.ListOptions{PerPage: 10}
	releases, _, err := client.Repositories.ListReleases(ctx, "kubernetes", "kubernetes", opt)
	if err != nil {
		panic(err) // is this really a good way?
	}
	minVersion := semver.New("1.8.0")
	allReleases := make([]*semver.Version, len(releases))
	for i, release := range releases {
		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}
		allReleases[i] = semver.New(versionString)
	}
	versionSlice := LatestVersions(allReleases, minVersion)

	fmt.Printf("latest versions of kubernetes/kubernetes: %s", versionSlice)
}
