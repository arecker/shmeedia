#!./venv/bin/python
import contextlib
import logging
import logging.config
import pathlib
import tempfile

import git

logger = logging.getLogger(__name__)
LOGGING_CONFIG = {
    'version': 1,
    'disable_existing_loggers': True,
    'formatters': {
        'standard': {
            'format': '%(asctime)s [%(levelname)s] %(name)s: %(message)s'
        },
    },
    'handlers': {
        'default': {
            'level': 'INFO',
            'formatter': 'standard',
            'class': 'logging.StreamHandler',
            'stream': 'ext://sys.stdout',  # Default is stderr
        },
    },
    'loggers': {
        '__main__': {
            'handlers': ['default'],
            'level': 'DEBUG',
            'propagate': False
        },
    }
}


def purge_output_dir():
    output_dir = pathlib.Path('output')
    assert output_dir.is_dir()
    output_dir_files = output_dir.glob('*.*')
    output_dir_files = filter(lambda f: not f.name.startswith('.'),
                              output_dir_files)
    output_dir_files = list(output_dir_files)
    map(lambda f: f.unlink(), output_dir_files)
    return len(output_dir_files)


@contextlib.contextmanager
def temp_git_repo():
    with tempfile.TemporaryDirectory() as temp:
        repo = git.Repo.init(temp, bare=True)
        logger.info('created temporary repo %s', repo)
        yield repo


def main():
    logging.config.dictConfig(LOGGING_CONFIG)

    num = purge_output_dir()
    logger.info('purged %d output file(s)', num)

    with temp_git_repo() as repo:
        pass

    shmeedia_bin = pathlib.Path('bin/shmeedia')
    assert shmeedia_bin.is_file()


if __name__ == '__main__':
    main()
