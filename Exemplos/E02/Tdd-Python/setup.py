from setuptools import setup, find_packages

setup(
    name='calculator',
    version='0.1.0',
    author='Seu Nome',
    author_email='seuemail@example.com',
    description='Uma breve descrição do seu módulo',
    long_description=open('README.md').read(),
    long_description_content_type='text/markdown',
    url='http://github.com/seuperfil/meu_modulo',
    packages=find_packages(exclude=('tests*',)),
    install_requires=[
        # Lista de dependências necessárias para o seu módulo
        # Exemplo: 'requests>=2.25.1',
    ],
    classifiers=[
        'Development Status :: 3 - Alpha',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.7',
        'Programming Language :: Python :: 3.8',
        'Programming Language :: Python :: 3.9',
        'Programming Language :: Python :: 3.10',
    ],
)
