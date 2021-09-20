run:
	uvicorn backend.main:app

run_reload:
	uvicorn backend.main:app --reload
