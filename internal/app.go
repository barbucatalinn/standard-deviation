package internal

import "strconv"

// App is the main struct of the application
type App struct {
	data []float64
}

// New creates a new app
func New(filename string) (*App, error) {
	app := new(App)

	// load data
	err := app.loadData(filename)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// loadData loads the csv data
func (app *App) loadData(filename string) error {
	// load CSV data
	csvData, err := readCSVFile(filename)
	if err != nil {
		return err
	}

	// format to float64
	for _, i := range csvData {
		if f, err := strconv.ParseFloat(i[0], 64); err == nil {
			app.data = append(app.data, f)
		}
	}

	return nil
}

// CalculateStandardDeviationV1 call the 'calculateStandardDeviationV1' function
// and returns the result
func (app *App) CalculateStandardDeviationV1() float64 {
	return calculateStandardDeviationV1(app.data)
}

// CalculateStandardDeviationV2 call the 'calculateStandardDeviationV2' function
// and returns the result
func (app *App) CalculateStandardDeviationV2() float64 {
	return calculateStandardDeviationV2(app.data)
}
