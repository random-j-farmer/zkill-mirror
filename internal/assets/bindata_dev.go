// Code generated by go-bindata.
// sources:
// static/pure-release-0.6.0/HISTORY.md
// static/pure-release-0.6.0/LICENSE.md
// static/pure-release-0.6.0/README.md
// static/pure-release-0.6.0/base-context-min.css
// static/pure-release-0.6.0/base-context.css
// static/pure-release-0.6.0/base-min.css
// static/pure-release-0.6.0/base.css
// static/pure-release-0.6.0/bower.json
// static/pure-release-0.6.0/buttons-core-min.css
// static/pure-release-0.6.0/buttons-core.css
// static/pure-release-0.6.0/buttons-min.css
// static/pure-release-0.6.0/buttons.css
// static/pure-release-0.6.0/forms-min.css
// static/pure-release-0.6.0/forms-nr-min.css
// static/pure-release-0.6.0/forms-nr.css
// static/pure-release-0.6.0/forms.css
// static/pure-release-0.6.0/grids-core-min.css
// static/pure-release-0.6.0/grids-core.css
// static/pure-release-0.6.0/grids-min.css
// static/pure-release-0.6.0/grids-responsive-min.css
// static/pure-release-0.6.0/grids-responsive-old-ie-min.css
// static/pure-release-0.6.0/grids-responsive-old-ie.css
// static/pure-release-0.6.0/grids-responsive.css
// static/pure-release-0.6.0/grids-units-min.css
// static/pure-release-0.6.0/grids-units.css
// static/pure-release-0.6.0/grids.css
// static/pure-release-0.6.0/menus-core-min.css
// static/pure-release-0.6.0/menus-core.css
// static/pure-release-0.6.0/menus-dropdown-min.css
// static/pure-release-0.6.0/menus-dropdown.css
// static/pure-release-0.6.0/menus-horizontal-min.css
// static/pure-release-0.6.0/menus-horizontal.css
// static/pure-release-0.6.0/menus-min.css
// static/pure-release-0.6.0/menus-scrollable-min.css
// static/pure-release-0.6.0/menus-scrollable.css
// static/pure-release-0.6.0/menus-skin-min.css
// static/pure-release-0.6.0/menus-skin.css
// static/pure-release-0.6.0/menus.css
// static/pure-release-0.6.0/pure-min.css
// static/pure-release-0.6.0/pure-nr-min.css
// static/pure-release-0.6.0/pure-nr.css
// static/pure-release-0.6.0/pure.css
// static/pure-release-0.6.0/tables-min.css
// static/pure-release-0.6.0/tables.css
// templates/touched.txt
// DO NOT EDIT!

// +build dev

package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// staticPureRelease060HistoryMd reads file data from disk. It returns an error on failure.
func staticPureRelease060HistoryMd() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/HISTORY.md")
	name := "static/pure-release-0.6.0/HISTORY.md"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060LicenseMd reads file data from disk. It returns an error on failure.
func staticPureRelease060LicenseMd() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/LICENSE.md")
	name := "static/pure-release-0.6.0/LICENSE.md"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060ReadmeMd reads file data from disk. It returns an error on failure.
func staticPureRelease060ReadmeMd() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/README.md")
	name := "static/pure-release-0.6.0/README.md"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060BaseContextMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060BaseContextMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/base-context-min.css")
	name := "static/pure-release-0.6.0/base-context-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060BaseContextCss reads file data from disk. It returns an error on failure.
func staticPureRelease060BaseContextCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/base-context.css")
	name := "static/pure-release-0.6.0/base-context.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060BaseMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060BaseMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/base-min.css")
	name := "static/pure-release-0.6.0/base-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060BaseCss reads file data from disk. It returns an error on failure.
func staticPureRelease060BaseCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/base.css")
	name := "static/pure-release-0.6.0/base.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060BowerJson reads file data from disk. It returns an error on failure.
func staticPureRelease060BowerJson() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/bower.json")
	name := "static/pure-release-0.6.0/bower.json"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060ButtonsCoreMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060ButtonsCoreMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/buttons-core-min.css")
	name := "static/pure-release-0.6.0/buttons-core-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060ButtonsCoreCss reads file data from disk. It returns an error on failure.
func staticPureRelease060ButtonsCoreCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/buttons-core.css")
	name := "static/pure-release-0.6.0/buttons-core.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060ButtonsMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060ButtonsMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/buttons-min.css")
	name := "static/pure-release-0.6.0/buttons-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060ButtonsCss reads file data from disk. It returns an error on failure.
func staticPureRelease060ButtonsCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/buttons.css")
	name := "static/pure-release-0.6.0/buttons.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060FormsMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060FormsMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/forms-min.css")
	name := "static/pure-release-0.6.0/forms-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060FormsNrMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060FormsNrMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/forms-nr-min.css")
	name := "static/pure-release-0.6.0/forms-nr-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060FormsNrCss reads file data from disk. It returns an error on failure.
func staticPureRelease060FormsNrCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/forms-nr.css")
	name := "static/pure-release-0.6.0/forms-nr.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060FormsCss reads file data from disk. It returns an error on failure.
func staticPureRelease060FormsCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/forms.css")
	name := "static/pure-release-0.6.0/forms.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060GridsCoreMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060GridsCoreMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/grids-core-min.css")
	name := "static/pure-release-0.6.0/grids-core-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060GridsCoreCss reads file data from disk. It returns an error on failure.
func staticPureRelease060GridsCoreCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/grids-core.css")
	name := "static/pure-release-0.6.0/grids-core.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060GridsMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060GridsMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/grids-min.css")
	name := "static/pure-release-0.6.0/grids-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060GridsResponsiveMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060GridsResponsiveMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/grids-responsive-min.css")
	name := "static/pure-release-0.6.0/grids-responsive-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060GridsResponsiveOldIeMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060GridsResponsiveOldIeMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/grids-responsive-old-ie-min.css")
	name := "static/pure-release-0.6.0/grids-responsive-old-ie-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060GridsResponsiveOldIeCss reads file data from disk. It returns an error on failure.
func staticPureRelease060GridsResponsiveOldIeCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/grids-responsive-old-ie.css")
	name := "static/pure-release-0.6.0/grids-responsive-old-ie.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060GridsResponsiveCss reads file data from disk. It returns an error on failure.
func staticPureRelease060GridsResponsiveCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/grids-responsive.css")
	name := "static/pure-release-0.6.0/grids-responsive.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060GridsUnitsMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060GridsUnitsMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/grids-units-min.css")
	name := "static/pure-release-0.6.0/grids-units-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060GridsUnitsCss reads file data from disk. It returns an error on failure.
func staticPureRelease060GridsUnitsCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/grids-units.css")
	name := "static/pure-release-0.6.0/grids-units.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060GridsCss reads file data from disk. It returns an error on failure.
func staticPureRelease060GridsCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/grids.css")
	name := "static/pure-release-0.6.0/grids.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusCoreMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusCoreMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-core-min.css")
	name := "static/pure-release-0.6.0/menus-core-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusCoreCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusCoreCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-core.css")
	name := "static/pure-release-0.6.0/menus-core.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusDropdownMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusDropdownMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-dropdown-min.css")
	name := "static/pure-release-0.6.0/menus-dropdown-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusDropdownCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusDropdownCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-dropdown.css")
	name := "static/pure-release-0.6.0/menus-dropdown.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusHorizontalMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusHorizontalMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-horizontal-min.css")
	name := "static/pure-release-0.6.0/menus-horizontal-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusHorizontalCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusHorizontalCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-horizontal.css")
	name := "static/pure-release-0.6.0/menus-horizontal.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-min.css")
	name := "static/pure-release-0.6.0/menus-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusScrollableMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusScrollableMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-scrollable-min.css")
	name := "static/pure-release-0.6.0/menus-scrollable-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusScrollableCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusScrollableCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-scrollable.css")
	name := "static/pure-release-0.6.0/menus-scrollable.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusSkinMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusSkinMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-skin-min.css")
	name := "static/pure-release-0.6.0/menus-skin-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusSkinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusSkinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus-skin.css")
	name := "static/pure-release-0.6.0/menus-skin.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060MenusCss reads file data from disk. It returns an error on failure.
func staticPureRelease060MenusCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/menus.css")
	name := "static/pure-release-0.6.0/menus.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060PureMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060PureMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/pure-min.css")
	name := "static/pure-release-0.6.0/pure-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060PureNrMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060PureNrMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/pure-nr-min.css")
	name := "static/pure-release-0.6.0/pure-nr-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060PureNrCss reads file data from disk. It returns an error on failure.
func staticPureRelease060PureNrCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/pure-nr.css")
	name := "static/pure-release-0.6.0/pure-nr.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060PureCss reads file data from disk. It returns an error on failure.
func staticPureRelease060PureCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/pure.css")
	name := "static/pure-release-0.6.0/pure.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060TablesMinCss reads file data from disk. It returns an error on failure.
func staticPureRelease060TablesMinCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/tables-min.css")
	name := "static/pure-release-0.6.0/tables-min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// staticPureRelease060TablesCss reads file data from disk. It returns an error on failure.
func staticPureRelease060TablesCss() (*asset, error) {
	path := filepath.Join(rootDir, "static/pure-release-0.6.0/tables.css")
	name := "static/pure-release-0.6.0/tables.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesTouchedTxt reads file data from disk. It returns an error on failure.
func templatesTouchedTxt() (*asset, error) {
	path := filepath.Join(rootDir, "templates/touched.txt")
	name := "templates/touched.txt"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"static/pure-release-0.6.0/HISTORY.md": staticPureRelease060HistoryMd,
	"static/pure-release-0.6.0/LICENSE.md": staticPureRelease060LicenseMd,
	"static/pure-release-0.6.0/README.md": staticPureRelease060ReadmeMd,
	"static/pure-release-0.6.0/base-context-min.css": staticPureRelease060BaseContextMinCss,
	"static/pure-release-0.6.0/base-context.css": staticPureRelease060BaseContextCss,
	"static/pure-release-0.6.0/base-min.css": staticPureRelease060BaseMinCss,
	"static/pure-release-0.6.0/base.css": staticPureRelease060BaseCss,
	"static/pure-release-0.6.0/bower.json": staticPureRelease060BowerJson,
	"static/pure-release-0.6.0/buttons-core-min.css": staticPureRelease060ButtonsCoreMinCss,
	"static/pure-release-0.6.0/buttons-core.css": staticPureRelease060ButtonsCoreCss,
	"static/pure-release-0.6.0/buttons-min.css": staticPureRelease060ButtonsMinCss,
	"static/pure-release-0.6.0/buttons.css": staticPureRelease060ButtonsCss,
	"static/pure-release-0.6.0/forms-min.css": staticPureRelease060FormsMinCss,
	"static/pure-release-0.6.0/forms-nr-min.css": staticPureRelease060FormsNrMinCss,
	"static/pure-release-0.6.0/forms-nr.css": staticPureRelease060FormsNrCss,
	"static/pure-release-0.6.0/forms.css": staticPureRelease060FormsCss,
	"static/pure-release-0.6.0/grids-core-min.css": staticPureRelease060GridsCoreMinCss,
	"static/pure-release-0.6.0/grids-core.css": staticPureRelease060GridsCoreCss,
	"static/pure-release-0.6.0/grids-min.css": staticPureRelease060GridsMinCss,
	"static/pure-release-0.6.0/grids-responsive-min.css": staticPureRelease060GridsResponsiveMinCss,
	"static/pure-release-0.6.0/grids-responsive-old-ie-min.css": staticPureRelease060GridsResponsiveOldIeMinCss,
	"static/pure-release-0.6.0/grids-responsive-old-ie.css": staticPureRelease060GridsResponsiveOldIeCss,
	"static/pure-release-0.6.0/grids-responsive.css": staticPureRelease060GridsResponsiveCss,
	"static/pure-release-0.6.0/grids-units-min.css": staticPureRelease060GridsUnitsMinCss,
	"static/pure-release-0.6.0/grids-units.css": staticPureRelease060GridsUnitsCss,
	"static/pure-release-0.6.0/grids.css": staticPureRelease060GridsCss,
	"static/pure-release-0.6.0/menus-core-min.css": staticPureRelease060MenusCoreMinCss,
	"static/pure-release-0.6.0/menus-core.css": staticPureRelease060MenusCoreCss,
	"static/pure-release-0.6.0/menus-dropdown-min.css": staticPureRelease060MenusDropdownMinCss,
	"static/pure-release-0.6.0/menus-dropdown.css": staticPureRelease060MenusDropdownCss,
	"static/pure-release-0.6.0/menus-horizontal-min.css": staticPureRelease060MenusHorizontalMinCss,
	"static/pure-release-0.6.0/menus-horizontal.css": staticPureRelease060MenusHorizontalCss,
	"static/pure-release-0.6.0/menus-min.css": staticPureRelease060MenusMinCss,
	"static/pure-release-0.6.0/menus-scrollable-min.css": staticPureRelease060MenusScrollableMinCss,
	"static/pure-release-0.6.0/menus-scrollable.css": staticPureRelease060MenusScrollableCss,
	"static/pure-release-0.6.0/menus-skin-min.css": staticPureRelease060MenusSkinMinCss,
	"static/pure-release-0.6.0/menus-skin.css": staticPureRelease060MenusSkinCss,
	"static/pure-release-0.6.0/menus.css": staticPureRelease060MenusCss,
	"static/pure-release-0.6.0/pure-min.css": staticPureRelease060PureMinCss,
	"static/pure-release-0.6.0/pure-nr-min.css": staticPureRelease060PureNrMinCss,
	"static/pure-release-0.6.0/pure-nr.css": staticPureRelease060PureNrCss,
	"static/pure-release-0.6.0/pure.css": staticPureRelease060PureCss,
	"static/pure-release-0.6.0/tables-min.css": staticPureRelease060TablesMinCss,
	"static/pure-release-0.6.0/tables.css": staticPureRelease060TablesCss,
	"templates/touched.txt": templatesTouchedTxt,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"static": &bintree{nil, map[string]*bintree{
		"pure-release-0.6.0": &bintree{nil, map[string]*bintree{
			"HISTORY.md": &bintree{staticPureRelease060HistoryMd, map[string]*bintree{}},
			"LICENSE.md": &bintree{staticPureRelease060LicenseMd, map[string]*bintree{}},
			"README.md": &bintree{staticPureRelease060ReadmeMd, map[string]*bintree{}},
			"base-context-min.css": &bintree{staticPureRelease060BaseContextMinCss, map[string]*bintree{}},
			"base-context.css": &bintree{staticPureRelease060BaseContextCss, map[string]*bintree{}},
			"base-min.css": &bintree{staticPureRelease060BaseMinCss, map[string]*bintree{}},
			"base.css": &bintree{staticPureRelease060BaseCss, map[string]*bintree{}},
			"bower.json": &bintree{staticPureRelease060BowerJson, map[string]*bintree{}},
			"buttons-core-min.css": &bintree{staticPureRelease060ButtonsCoreMinCss, map[string]*bintree{}},
			"buttons-core.css": &bintree{staticPureRelease060ButtonsCoreCss, map[string]*bintree{}},
			"buttons-min.css": &bintree{staticPureRelease060ButtonsMinCss, map[string]*bintree{}},
			"buttons.css": &bintree{staticPureRelease060ButtonsCss, map[string]*bintree{}},
			"forms-min.css": &bintree{staticPureRelease060FormsMinCss, map[string]*bintree{}},
			"forms-nr-min.css": &bintree{staticPureRelease060FormsNrMinCss, map[string]*bintree{}},
			"forms-nr.css": &bintree{staticPureRelease060FormsNrCss, map[string]*bintree{}},
			"forms.css": &bintree{staticPureRelease060FormsCss, map[string]*bintree{}},
			"grids-core-min.css": &bintree{staticPureRelease060GridsCoreMinCss, map[string]*bintree{}},
			"grids-core.css": &bintree{staticPureRelease060GridsCoreCss, map[string]*bintree{}},
			"grids-min.css": &bintree{staticPureRelease060GridsMinCss, map[string]*bintree{}},
			"grids-responsive-min.css": &bintree{staticPureRelease060GridsResponsiveMinCss, map[string]*bintree{}},
			"grids-responsive-old-ie-min.css": &bintree{staticPureRelease060GridsResponsiveOldIeMinCss, map[string]*bintree{}},
			"grids-responsive-old-ie.css": &bintree{staticPureRelease060GridsResponsiveOldIeCss, map[string]*bintree{}},
			"grids-responsive.css": &bintree{staticPureRelease060GridsResponsiveCss, map[string]*bintree{}},
			"grids-units-min.css": &bintree{staticPureRelease060GridsUnitsMinCss, map[string]*bintree{}},
			"grids-units.css": &bintree{staticPureRelease060GridsUnitsCss, map[string]*bintree{}},
			"grids.css": &bintree{staticPureRelease060GridsCss, map[string]*bintree{}},
			"menus-core-min.css": &bintree{staticPureRelease060MenusCoreMinCss, map[string]*bintree{}},
			"menus-core.css": &bintree{staticPureRelease060MenusCoreCss, map[string]*bintree{}},
			"menus-dropdown-min.css": &bintree{staticPureRelease060MenusDropdownMinCss, map[string]*bintree{}},
			"menus-dropdown.css": &bintree{staticPureRelease060MenusDropdownCss, map[string]*bintree{}},
			"menus-horizontal-min.css": &bintree{staticPureRelease060MenusHorizontalMinCss, map[string]*bintree{}},
			"menus-horizontal.css": &bintree{staticPureRelease060MenusHorizontalCss, map[string]*bintree{}},
			"menus-min.css": &bintree{staticPureRelease060MenusMinCss, map[string]*bintree{}},
			"menus-scrollable-min.css": &bintree{staticPureRelease060MenusScrollableMinCss, map[string]*bintree{}},
			"menus-scrollable.css": &bintree{staticPureRelease060MenusScrollableCss, map[string]*bintree{}},
			"menus-skin-min.css": &bintree{staticPureRelease060MenusSkinMinCss, map[string]*bintree{}},
			"menus-skin.css": &bintree{staticPureRelease060MenusSkinCss, map[string]*bintree{}},
			"menus.css": &bintree{staticPureRelease060MenusCss, map[string]*bintree{}},
			"pure-min.css": &bintree{staticPureRelease060PureMinCss, map[string]*bintree{}},
			"pure-nr-min.css": &bintree{staticPureRelease060PureNrMinCss, map[string]*bintree{}},
			"pure-nr.css": &bintree{staticPureRelease060PureNrCss, map[string]*bintree{}},
			"pure.css": &bintree{staticPureRelease060PureCss, map[string]*bintree{}},
			"tables-min.css": &bintree{staticPureRelease060TablesMinCss, map[string]*bintree{}},
			"tables.css": &bintree{staticPureRelease060TablesCss, map[string]*bintree{}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"touched.txt": &bintree{templatesTouchedTxt, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

