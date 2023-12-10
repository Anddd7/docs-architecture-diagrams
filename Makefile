translate:
	@cd cmds/translate && go run .

translate_all_en:
	@cd cmds/translate && go run . ../.. en
