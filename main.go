/*
Copyright 2019 Marilyn Chua <marilyn.c@privasec.com>
Purpose: Automate Reconnaissance Process
Date:    August 2019 

Example of Usage:
$ go build main.go
$ sudo ./main
*/

package main
 
import (
    "fmt"
    // "plugin"
    // "path/filepath"
    // "../golang/core/config"
    "../golang/core/util"
    "../golang/core/processor"
)
 
func main() {
    /* 
    Print banner and compile plugins 
    */
    util.Initialize()

    /* 
    [Future Enhancement]
    Set scan configurations
    Options - Mode: Quiet, Quick, Comprehensive
    */
    options := "10.10.10.101"
 
    targets := make(map[string]map[string][]string)
    targets = processor.SetTargets(options)

    fmt.Println(targets)

    /* 
    Load plugins from 'core/plugins' directory
    Pre-requisite: Plugins must be compiled (*.so)
    */
    // all_plugins, err := filepath.Glob("core/plugins/*.so")
    // if err != nil {
    //     panic(err)
    // }
 
    /* 
    Load plugins by stages
    Function: Current plugins are load by alphabetical orders
    */
    // for _, stage := range (config.C_PHASE) {
    //     for _, filename := range (all_plugins) {
    //         p, err := plugin.Open(filename)
    //         if err != nil {
    //             panic(err)
    //         }
             
    //         function, err := p.Lookup(stage)
    //         if err != nil {
    //             continue
    //         }
     
    //         _function, ok := function.(func(map[string]string) *map[string]string)
    //         if !ok {
    //             fmt.Printf("Plugin has no '%s(map[string]string) map[string]string' function\n", stage)
    //             continue
    //         }
     
    //         // scan_function(options)
    //         scanned := _function(options)
    //         fmt.Println(filename, scanned)
    //     }
    // }
    
}