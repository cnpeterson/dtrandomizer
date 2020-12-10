package main

import (
    "log"
    "io/ioutil"
    "math/rand"
    "os/exec"
    "time"
    "gopkg.in/yaml.v2"
    "flag"
    "path/filepath"
    "os"
)

var filename = flag.String("config", "conf.yaml", "Location of the config file")
var images []string

func main() {

    var c conf
    c.getConf()
    err := filepath.Walk(c.Dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            ext := filepath.Ext(path)
            for _, e := range c.Ext {
                if e == ext {
                    images = append(images, path)
                    continue
                }
            }
        }

        return nil
    })

    if err != nil {
        log.Println(err)
    }

    num := numberPicker(len(images))

    f := images[num]
    cmd := exec.Command("/usr/bin/feh", "--bg-fill", f)
    _, err = cmd.Output()
    if err != nil {
        log.Fatal(err)
    }
}

func numberPicker(max int) int {
    rand.Seed(time.Now().UnixNano())
    num := rand.Intn(max)
    return num
}

type conf struct {
    Dir string `yaml:"dir"`
    Ext []string `yaml:"fileExtensions"`
}

func (c *conf) getConf() *conf {
    flag.Parse()
    log.Println(*filename)
    yamlFile, err := ioutil.ReadFile(*filename)
    if err !=  nil {
        log.Fatal(err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatal(err)
    }

    return c
}
