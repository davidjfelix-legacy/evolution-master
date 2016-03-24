extern crate cpython;
extern crate ansi_term;

use ansi_term::Colour::Green;
use std::fs;
use std::path::Path;
use std::process::{Command, Stdio};
use cpython::Python;
use cpython::ObjectProtocol;

fn main() {
    let directory: &Path;

    if cfg!(unix) {
        directory  = Path::new("/opt/evolution-master");
    } else if cfg!(windows) {
        directory = Path::new("C:/Program Data/evolution-master");
    } else {
        directory = Path::new("./evolution-master");
    } 

    print_header();
    get_ubuntu_deps();
    build_python();

    println!("Spin sequences: {}", Green.paint("initiated"));
    match fs::create_dir_all(directory) {
        Err(_) => println!("failed"),
        _ => println!("Worked")
    }

    let gil = Python::acquire_gil();
    let py = gil.python();

    let sys = py.import("sys").unwrap();
    let version: String = sys.get(py, "version").unwrap().extract(py).unwrap();

    let os = py.import("os").unwrap();
    let getenv = os.get(py, "getenv").unwrap();
    let user: String = getenv.call(py, ("USER",), None).unwrap().extract(py).unwrap();

    println!("Hello {}, I'm Python {}", user, version);
}

fn print_header() {
    println!("Evolution Master: {}", Green.paint("awoken"));
    println!("Discovering environment");

}

fn get_ubuntu_deps() {
    Command::new("apt-get")
        .arg("update")
        .stdout(Stdio::inherit())
        .output();
    Command::new("apt-get")
        .arg("-y")
        .arg("install")
        .stdout(Stdio::inherit())
        .arg("git")
        .output();
}

fn build_python() {
    // FIXME: stop panicking here. trap errors
    let output = Command::new("mktemp")
        .arg("-d")
        .arg("evolution-master.python.XXXXXX")
        .arg("--tmpdir")
        .output()
        .unwrap();
    let temp_dir = String::from_utf8(output.stdout).unwrap();
    let tar = temp_dir.clone() + "Python-3.5.0.tar.xz";
    Command::new("wget")
        .arg("https://www.python.org/ftp/python/3.5.0/Python-3.5.0.tar.xz")
        .arg("-O")
        .arg(&tar) 
        .stdout(Stdio::inherit())
        .output();
    Command::new("tar")
        .arg("xvf")
        .arg(&tar)
        .stdout(Stdio::inherit())
        .output();
    Command::new(temp_dir.clone() + "/Python-3.5.0/configure")
        .arg("--prefix=/opt/evolution-master/python")
        .stdout(Stdio::inherit())
        .output();
    let makefile_flag = "-f".to_string() + &temp_dir.clone() + "/Python-3.5.0/Makefile";
    Command::new("make")
        .arg("-j16")
        .arg(&makefile_flag)
        .stdout(Stdio::inherit())
        .output();
    Command::new("make")
        .arg("install")
        .arg(&makefile_flag)
        .stdout(Stdio::inherit())
        .output();
}
