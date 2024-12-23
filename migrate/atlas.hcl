data "external_schema" "gorm" {
  program = ["go", "run", "./main.go"]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://mysql/8/dev"
  migration {
    dir = "file://mysql"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}