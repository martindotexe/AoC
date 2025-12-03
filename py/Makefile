run:
	@for dir in day*/; do \
		if [ -f "$$dir/main.by" ]; then \
			uv run bython -c -o "$$dir/main.py" "$$dir/main.by"; \
		fi \
	done
	uv run bython main.by

clean:
	rm -f main.py day*/main.py
