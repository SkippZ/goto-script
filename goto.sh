goto() {
    local out rc
    # capture stdout (suppress program stderr; remove 2>/dev/null if you want to see errors)
    out="$("${HOME}/projects/tools/goto/goto" "$@" 2>/dev/null)" || rc=$?

    if [ -z "$out" ]; then
        printf 'goto_eval: program produced no output\n' >&2
        return ${rc:-1}
    fi

    # Basic sanity check: output should start with "cd " (cd + whitespace)
    if [[ $out == cd\ * ]]; then
        # run the printed command in the current shell
        eval -- "$out" || {
            printf 'goto_eval: eval failed\n' >&2
                    return 1
                }
        else
            printf 'goto_eval: unexpected output (not starting with "cd "): %s\n' "$out" >&2
            return 1
    fi
}
