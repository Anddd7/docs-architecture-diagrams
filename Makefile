build_translate:
	@mkdir -p bin
	@cd cmds/translate && go build -o ../../bin/translate . 

translate_en: build_translate
	@bin/translate diagrams en

# bin/translate <path> <language>