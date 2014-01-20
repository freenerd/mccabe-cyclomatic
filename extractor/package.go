package extractor

import (
	"fmt"
	"go/build"
	"os"
	"path"
)

func PackageImportCalls(pkgName string) (int64, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return -1, fmt.Errorf("failed to get working dir: %s", err)
	}

	complexity, err := processPackage(cwd, pkgName)

	if err != nil {
		return -1, fmt.Errorf("failed to process package: %s", err)
	}

	return complexity, nil
}

func processPackage(root, pkgName string) (int64, error) {
	// read package
	pkg, err := build.Import(pkgName, root, 0)
	if err != nil {
		return -1, fmt.Errorf("failed to import package %s: %s", pkgName, err)
	}

	// Don't worry about dependencies for stdlib packages
	if pkg.Goroot {
		return 0, nil
	}

	var packageComplexity int64 = 0

	// analyze each file in package, merge results
	for _, file := range pkg.GoFiles {
		fileComplexity, err := FileImportCalls(path.Join(pkg.Dir, file), pkg)
		if err != nil {
			return -1, fmt.Errorf("failed in file %s: %s", file, err)
		}

		packageComplexity = packageComplexity + fileComplexity
	}

	return packageComplexity, nil
}
