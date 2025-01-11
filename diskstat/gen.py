from pathlib import Path

CURRENT = Path(__file__).resolve().parent
gen_dir = CURRENT.parent.joinpath("diskstat_script")
script = gen_dir.joinpath("Diskstat.ahk")


def gen():
    gen_dir.mkdir(exist_ok=True)

    content = CURRENT.joinpath("Diskstat.template.ahk").read_text(encoding="utf8")
    content = content.replace("{resources}", str(CURRENT))

    script.write_text(content, encoding="utf8")
