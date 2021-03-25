package cmd

import (
	"strconv"

	"github.com/klec/demo/pkg/log"
	"github.com/sjwhitworth/golearn/base"
	"github.com/urfave/cli"

	"github.com/sajari/regression"
	"github.com/sjwhitworth/golearn/linear_models"
)

var (
	Linear = cli.Command{
		Name:        "linear",
		Usage:       "demo linear <option>",
		Description: "run demo linear",
		Subcommands: []cli.Command{
			subCmdLinear,
			subCmdLinearModel,
			subCmdLearn,
		},
	}

	subCmdLinear = cli.Command{
		Name:        "regression",
		Usage:       "demo for linear regression",
		Description: "run demo for linear regression",
		Flags: []cli.Flag{
			floatFlag("float, f", 8.6, "Predict value"),
		},
		Action: runLinearRegression,
	}

	subCmdLinearModel = cli.Command{
		Name:        "model",
		Usage:       "demo for linear model",
		Description: "run demo for linear model",
		Action:      runLinearModel,
	}

	subCmdLearn = cli.Command{
		Name:        "learn",
		Usage:       "demo for linear model learn",
		Description: "run demo for linear model learn",
		Action:      runLinearModelLearn,
	}
)

func runLinearRegression(ctx *cli.Context) error {
	r := new(regression.Regression)
	r.SetObserved("Gas value")
	r.SetVar(0, "k")
	r.Train(
		regression.DataPoint(5.0, []float64{2.5}),
		regression.DataPoint(8.0, []float64{4.0}),
		regression.DataPoint(9.6, []float64{4.8}),
		regression.DataPoint(12.0, []float64{6.0}),
		regression.DataPoint(12.0, []float64{6.0}),
		regression.DataPoint(13.0, []float64{6.5}),
		regression.DataPoint(14.0, []float64{7.0}),
		regression.DataPoint(14.0, []float64{7.0}),
		regression.DataPoint(15.0, []float64{7.5}),
		regression.DataPoint(18.0, []float64{9.0}),
		regression.DataPoint(20.0, []float64{10.0}),
		regression.DataPoint(21.0, []float64{10.5}),
		regression.DataPoint(24.8, []float64{12.4}),
		regression.DataPoint(25.0, []float64{12.5}),
		regression.DataPoint(25.8, []float64{17.9}),
		regression.DataPoint(26.0, []float64{13.0}),
		regression.DataPoint(28.0, []float64{14.0}),
		regression.DataPoint(35.0, []float64{17.5}),
		regression.DataPoint(36.2, []float64{18.1}),
		regression.DataPoint(40.0, []float64{20.0}),
	)
	r.Run()
	log.Info("Regression formula:%v", r.Formula)

	prediction, _ := r.Predict([]float64{ctx.Float64("f")})
	log.Info("Regression prediction:%v", prediction)
	return nil
}

func runLinearModel(ctx *cli.Context) error {
	lr := linear_models.NewLinearRegression()
	//lr, err := linear_models.NewLogisticRegression("l2", 1, 0.5)
	//if err != nil {
	//	log.Error("Init NewLogisticRegression err:%s", err)
	//	return err
	//}

	train, err := base.ParseCSVToInstances("./internal/linear/train.csv", true)
	if err != nil {
		log.Error("Train error:%s", err)
		return nil
	}
	test, err := base.ParseCSVToInstances("./internal/linear/test.csv", true)
	if err != nil {
		log.Error("Test error:%s", err)
		return nil
	}

	err = lr.Fit(train)
	if err != nil {
		log.Error("Fit error:%s", err)
		return nil
	}

	prediction, _ := lr.Predict(test)
	_, rows := prediction.Size()
	for i := 0; i < rows; i++ {
		expected, _ := strconv.ParseFloat(base.GetClass(test, i), 64)
		actual, _ := strconv.ParseFloat(base.GetClass(prediction, i), 64)
		log.Info("Input:%s, expected:%0.3f, actual:%0.3f, result:%t", test.RowString(i), expected, actual, expected == actual)
	}

	return nil
}

func runLinearModelLearn(ctx *cli.Context) error {
	demo, err := base.ParseCSVToInstances("./internal/linear/demo.csv", true)
	if err != nil {
		log.Error("Demo error:%s", err)
		return nil
	}
	log.Info("Demo:\r\v%+v", demo)
	//log.Info("Demo string:%s", demo.RowString(0))
	//x, y := demo.Size()
	//log.Info("Demo size:%d, %d", x, y)

	newInst := base.NewDenseInstances()

	attrs := make([]base.Attribute, 3)
	attrs[0] = base.NewFloatAttribute("input1")
	attrs[1] = base.NewFloatAttribute("input2")
	attrs[2] = base.NewFloatAttribute("output")
	newSpecs := make([]base.AttributeSpec, 3)
	newSpecs[0] = newInst.AddAttribute(attrs[0])
	newSpecs[1] = newInst.AddAttribute(attrs[1])
	newSpecs[2] = newInst.AddAttribute(attrs[2])

	err = newInst.Extend(1)
	if err != nil {
		log.Error("Instance extend err:%s", err)
		return err
	}

	newInst.Set(newSpecs[0], 0, newSpecs[0].GetAttribute().GetSysValFromString("100.0"))
	newInst.Set(newSpecs[1], 0, newSpecs[1].GetAttribute().GetSysValFromString("200.0"))
	newInst.Set(newSpecs[2], 0, newSpecs[2].GetAttribute().GetSysValFromString("300.0"))
	log.Info("newInst:\r\v%+v", newInst)
	return nil
}
