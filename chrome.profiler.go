package chrome

import (
	"encoding/json"

	profiler "github.com/mkenney/go-chrome/profiler"
	"github.com/mkenney/go-chrome/protocol"

	log "github.com/Sirupsen/logrus"
)

/*
Profiler - https://chromedevtools.github.io/devtools-protocol/tot/Profiler/
Provides various functionality related to drawing atop the inspected page. EXPERIMENTAL
*/
type Profiler struct{}

/*
Disable disables profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-disable
*/
func (Profiler) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enable
*/
func (Profiler) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetBestEffortCoverage collects coverage data for the current isolate. The coverage data may be
incomplete due to garbage collection.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getBestEffortCoverage
*/
func (Profiler) GetBestEffortCoverage(
	socket *Socket,
) (profiler.GetBestEffortCoverageResult, error) {
	command := &protocol.Command{
		Method: "Profiler.getBestEffortCoverage",
	}
	result := profiler.GetBestEffortCoverageResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
SetSamplingInterval changes CPU profiler sampling interval. Must be called before CPU profiles
recording started.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-setSamplingInterval
*/
func (Profiler) SetSamplingInterval(
	socket *Socket,
	params *profiler.SetSamplingIntervalParams,
) error {
	command := &protocol.Command{
		Method: "Profiler.setSamplingInterval",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Start starts profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-start
*/
func (Profiler) Start(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.start",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartPreciseCoverage enable precise code coverage. Coverage data for JavaScript executed before
enabling precise code coverage may be incomplete. Enabling prevents running optimized code and
resets execution counters.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startPreciseCoverage
*/
func (Profiler) StartPreciseCoverage(
	socket *Socket,
	params *profiler.StartPreciseCoverageParams,
) error {
	command := &protocol.Command{
		Method: "Profiler.startPreciseCoverage",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StartTypeProfile enables type profile. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startTypeProfile
*/
func (Profiler) StartTypeProfile(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.startTypeProfile",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Stop stops profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stop
*/
func (Profiler) Stop(
	socket *Socket,
) (profiler.StopResult, error) {
	command := &protocol.Command{
		Method: "Profiler.stop",
	}
	result := profiler.StopResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
StopPreciseCoverage disable precise code coverage. Disabling releases unnecessary execution count
records and allows executing optimized code.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopPreciseCoverage
*/
func (Profiler) StopPreciseCoverage(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.stopPreciseCoverage",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
StopTypeProfile disables type profile. Disabling releases type profile data collected so far.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopTypeProfile
*/
func (Profiler) StopTypeProfile(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "Profiler.stopTypeProfile",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
TakePreciseCoverage collects coverage data for the current isolate, and resets execution counters.
Precise code coverage needs to have started.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takePreciseCoverage
*/
func (Profiler) TakePreciseCoverage(
	socket *Socket,
) (profiler.TakePreciseCoverageResult, error) {
	command := &protocol.Command{
		Method: "Profiler.takePreciseCoverage",
	}
	result := profiler.TakePreciseCoverageResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
TakeTypeProfile collect type profile. EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takeTypeProfile
*/
func (Profiler) TakeTypeProfile(
	socket *Socket,
) (profiler.TakeTypeProfileResult, error) {
	command := &protocol.Command{
		Method: "Profiler.takeTypeProfile",
	}
	result := profiler.TakeTypeProfileResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
OnConsoleProfileFinished adds a handler to the Profiler.consoleProfileFinished event.
Profiler.consoleProfileFinished fires when profile recording finishes.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileFinished
*/
func (Profiler) OnConsoleProfileFinished(
	socket *Socket,
	callback func(event *profiler.ConsoleProfileFinishedEvent),
) {
	handler := protocol.NewEventHandler(
		"Profiler.consoleProfileFinished",
		func(name string, params []byte) {
			event := &profiler.ConsoleProfileFinishedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnConsoleProfileStarted adds a handler to the Profiler.consoleProfileStarted event.
Profiler.consoleProfileStarted fires when new profile recording is started using console.profile()
call.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileStarted
*/
func (Profiler) OnConsoleProfileStarted(
	socket *Socket,
	callback func(event *profiler.ConsoleProfileStartedEvent),
) {
	handler := protocol.NewEventHandler(
		"Profiler.consoleProfileStarted",
		func(name string, params []byte) {
			event := &profiler.ConsoleProfileStartedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
