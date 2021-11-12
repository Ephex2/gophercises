package urlshort

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx"
	yml "gopkg.in/yaml.v2"
)

type innerConfig struct {
	Path string
	Url  string
}

type config struct {
	Pairs []innerConfig
}

func parseYaml(y []byte) (out config, err error) {
	err = yml.Unmarshal(y, &out.Pairs)
	return
}

func parseJson(j []byte) (out config, err error) {
	err = json.Unmarshal(j, &out)
	return
}

func getSqlPair(path string) (out config, err error) {
	conf, err := pgx.ParseConnectionString(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Error parsing database connection configuration from environment variable (DATABASE_URL). Error: " + err.Error())
	}

	conn, err := pgx.Connect(conf)
	if err != nil {
		fmt.Printf("Error connecting with connection info :%v", conn)
		panic(err)
	}
	defer conn.Close()

	var pair innerConfig
	err = conn.QueryRow("SELECT path, url FROM public.pairs where path = $1;", path).Scan(&pair.Path, &pair.Url)
	if err != nil {
		return
	} else if pair.Path == "" {
		err = errors.New("path not found")
	}

	out.Pairs = append(out.Pairs, pair)
	return
}

func buildMap(in config) map[string]string {
	out := make(map[string]string)

	for _, pair := range in.Pairs {
		out[pair.Path] = pair.Url
	}

	return out
}
