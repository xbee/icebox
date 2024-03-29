/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package flogging

import (
	//"github.com/op/go-logging"
	"github.com/cheapRoc/grpc-zerolog"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/grpclog"
)

const GRPCModuleID = "grpc"

func initgrpclogger() {
	glogger := MustGetLoggerWithDefaultLevel(GRPCModuleID)
	grpclog.SetLoggerV2(grpczerolog.New(*glogger))
}

// grpclogger implements the standard Go logging interface and wraps the
// logger provided by the flogging package.  This is required in order to
// replace the default log used by the grpclog package.
type grpclogger struct {
	logger *zerolog.Logger
}

func (g *grpclogger) Fatal(args ...interface{}) {
	g.logger.Fatal().Msgf("", args...)
}

func (g *grpclogger) Fatalf(format string, args ...interface{}) {
	g.logger.Fatal().Msgf(format, args...)
}

func (g *grpclogger) Fatalln(args ...interface{}) {
	g.logger.Fatal().Msgf("", args...)
}

// NOTE: grpclog does not support leveled logs so for now use DEBUG
func (g *grpclogger) Print(args ...interface{}) {
	g.logger.Debug().Msgf("", args...)
}

func (g *grpclogger) Printf(format string, args ...interface{}) {
	g.logger.Debug().Msgf(format, args...)
}

func (g *grpclogger) Println(args ...interface{}) {
	g.logger.Debug().Msgf("", args...)
}
