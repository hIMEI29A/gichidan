# Changelog - gichidan

### 0.1.0

__Changes__

- README.md updated
  CHANGELOG.md created
  Makefile removed
  ShortInfo option added to CLI
  All features: search requests, html parsing, results collecting, CLI, info output work fine.
- Save to file option fixed
  Subcommands removed, now CLI with flags only
- README.md added
- All goroutines starts in main() only
  Main() with infinite loop for {} now
  Visited links control in main() too
  GetTotal number of hosts added
  When len(parsedHosts) == Total, infinite for {} breaks
  Now here is no recursion in spider.Crawl()
- new file:   .gitignore
  new file:   AUTHORS
  new file:   LICENSE
  new file:   Makefile
  new file:   data.go
  new file:   glide.lock
  new file:   glide.yaml
  new file:   main.go
  new file:   parser.go



__Contributors__

- hIMEI

Released by hIMEI, Sun 21 Jan 2018 -
[see the diff](https://github.com/<no value>/gichidan/compare/e57581e8c548fee66ffbff1b7dea693ee27a7b2d...0.1.0#diff)
______________


