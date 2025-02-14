U
    rܢ^z*  �                   @   s�   d dl mZmZmZmZmZ d dlmZ d dlZd dl	Z	ddl
mZmZmZmZ d dl mZ d dlmZ G dd	� d	e�ZG d
d� de�ZG dd� de�ZG dd� de�ZG dd� de�ZdS )�    )�TType�TMessageType�TFrozenDict�
TException�TApplicationException)�TProtocolExceptionN�   )�LiffViewRequest�LiffException�LiffViewResponse�RevokeTokenRequest)�
TProcessor)�
TTransportc                   @   sj   e Zd ZdZddejdedgdffZddd�Zdd� Z	d	d
� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�issueLiffView_args�$
    Attributes:
     - request
    Nr   �requestc                 C   s
   || _ d S �N�r   ��selfr   � r   �//storage/sdcard0/bot1/flex/curve/LiffService.py�__init__   s    zissueLiffView_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n
|�|� |��  q6|��  | ��  d S �Nr   )�_fast_decode�thrift_spec�	__class__�readStructBegin�readFieldBeginr   �STOP�STRUCTr	   r   �read�skip�readFieldEnd�readStructEnd�validate�r   �iprot�fname�ftype�fidr   r   r   r!      s     



zissueLiffView_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S | ��  |�d� | jd k	rt|�dt	j
d� | j�|� |��  |��  |��  d S )Nr   r   r   )�_fast_encoder   �trans�writer   r%   �writeStructBeginr   �writeFieldBeginr   r    �writeFieldEnd�writeFieldStop�writeStructEnd�r   �oprotr   r   r   r-   .   s    

zissueLiffView_args.writec                 C   s   d S r   r   �r   r   r   r   r%   ;   s    zissueLiffView_args.validatec                 C   s   d}|d t t| j��A }|S �N�   �   ��hash�make_hashabler   �r   �valuer   r   r   �__hash__>   s    zissueLiffView_args.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS �z%s=%rr   ��.0�keyr=   r   r   r   �
<listcomp>D   s   �z/issueLiffView_args.__repr__.<locals>.<listcomp>�%s(%s)�, ��__dict__�itemsr   �__name__�join�r   �Lr   r   r   �__repr__C   s    �zissueLiffView_args.__repr__c                 C   s   t || j�o| j|jkS r   ��
isinstancer   rG   �r   �otherr   r   r   �__eq__H   s    zissueLiffView_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   �__ne__K   s    zissueLiffView_args.__ne__)N)rI   �
__module__�__qualname__�__doc__r   r    r	   r   r   r!   r-   r%   r>   rM   rR   rS   r   r   r   r   r      s   �
r   c                   @   sz   e Zd ZdZdejdedgdfdejdedgdffZddd�Z	d	d
� Z
dd� Zdd� Zdd� Zdd� Zdd� Zdd� ZdS )�issueLiffView_resultz-
    Attributes:
     - success
     - e
    r   �successNr   �ec                 C   s   || _ || _d S r   )rX   rY   )r   rX   rY   r   r   r   r   Y   s    zissueLiffView_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n>|dkr�|tjkr�t� | _| j�
|� q�|�|� n
|�|� |��  q6|��  | ��  d S )Nr   r   )r   r   r   r   r   r   r   r    r   rX   r!   r"   r
   rY   r#   r$   r%   r&   r   r   r   r!   ]   s*    




zissueLiffView_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S | ��  |�d� | jd k	rt|�dt	j
d� | j�|� |��  | jd k	r�|�dt	j
d� | j�|� |��  |��  |��  d S )NrW   rX   r   rY   r   )r+   r   r,   r-   r   r%   r.   rX   r/   r   r    r0   rY   r1   r2   r3   r   r   r   r-   x   s    


zissueLiffView_result.writec                 C   s   d S r   r   r5   r   r   r   r%   �   s    zissueLiffView_result.validatec                 C   s4   d}|d t t| j��A }|d t t| j��A }|S r6   )r:   r;   rX   rY   r<   r   r   r   r>   �   s    zissueLiffView_result.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rC   �   s   �z1issueLiffView_result.__repr__.<locals>.<listcomp>rD   rE   rF   rK   r   r   r   rM   �   s    �zissueLiffView_result.__repr__c                 C   s   t || j�o| j|jkS r   rN   rP   r   r   r   rR   �   s    zissueLiffView_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �   s    zissueLiffView_result.__ne__)NN)rI   rT   rU   rV   r   r    r   r
   r   r   r!   r-   r%   r>   rM   rR   rS   r   r   r   r   rW   N   s   �
rW   c                   @   sj   e Zd ZdZddejdedgdffZddd�Zdd� Z	d	d
� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�revokeToken_argsr   Nr   r   c                 C   s
   || _ d S r   r   r   r   r   r   r   �   s    zrevokeToken_args.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n
|�|� |��  q6|��  d S r   )r   r   r   r   r   r   r   r    r   r   r!   r"   r#   r$   r&   r   r   r   r!   �   s    



zrevokeToken_args.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )NrZ   r   r   )r+   r   r,   r-   r   r.   r   r/   r   r    r0   r1   r2   r3   r   r   r   r-   �   s    

zrevokeToken_args.writec                 C   s   d S r   r   r5   r   r   r   r%   �   s    zrevokeToken_args.validatec                 C   s   d}|d t t| j��A }|S r6   r9   r<   r   r   r   r>   �   s    zrevokeToken_args.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rC   �   s   �z-revokeToken_args.__repr__.<locals>.<listcomp>rD   rE   rF   rK   r   r   r   rM   �   s    �zrevokeToken_args.__repr__c                 C   s   t || j�o| j|jkS r   rN   rP   r   r   r   rR   �   s    zrevokeToken_args.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS   �   s    zrevokeToken_args.__ne__)N)rI   rT   rU   rV   r   r    r   r   r   r!   r-   r%   r>   rM   rR   rS   r   r   r   r   rZ   �   s   �
rZ   c                   @   sj   e Zd ZdZddejdedgdffZddd�Zdd� Z	d	d
� Z
dd� Zdd� Zdd� Zdd� Zdd� ZdS )�revokeToken_resultz
    Attributes:
     - e
    Nr   rY   c                 C   s
   || _ d S r   )rY   )r   rY   r   r   r   r   �   s    zrevokeToken_result.__init__c                 C   s�   |j d k	r.| jd k	r.|� | || j| jg� d S |��  |�� \}}}|tjkrPq�|dkr�|tjkrxt� | _	| j	�
|� q�|�|� n
|�|� |��  q6|��  d S r   )r   r   r   r   r   r   r   r    r
   rY   r!   r"   r#   r$   r&   r   r   r   r!   �   s    



zrevokeToken_result.readc                 C   s�   |j d k	r4| jd k	r4|j�|� | | j| jg�� d S |�d� | jd k	rl|�dtj	d� | j�|� |�
�  |��  |��  d S )Nr[   rY   r   )r+   r   r,   r-   r   r.   rY   r/   r   r    r0   r1   r2   r3   r   r   r   r-   �   s    

zrevokeToken_result.writec                 C   s   d S r   r   r5   r   r   r   r%   
  s    zrevokeToken_result.validatec                 C   s   d}|d t t| j��A }|S r6   )r:   r;   rY   r<   r   r   r   r>     s    zrevokeToken_result.__hash__c                 C   s*   dd� | j �� D �}d| jjd�|�f S )Nc                 S   s   g | ]\}}d ||f �qS r?   r   r@   r   r   r   rC     s   �z/revokeToken_result.__repr__.<locals>.<listcomp>rD   rE   rF   rK   r   r   r   rM     s    �zrevokeToken_result.__repr__c                 C   s   t || j�o| j|jkS r   rN   rP   r   r   r   rR     s    zrevokeToken_result.__eq__c                 C   s
   | |k S r   r   rP   r   r   r   rS     s    zrevokeToken_result.__ne__)N)rI   rT   rU   rV   r   r    r
   r   r   r!   r-   r%   r>   rM   rR   rS   r   r   r   r   r[   �   s   �
r[   c                   @   s&   e Zd Zddd�Zdd� Zdd� ZdS )	�ClientNc                 C   s$   | | _ | _|d k	r|| _d| _d S )Nr   )�_iprot�_oprot�_seqid)r   r'   r4   r   r   r   r     s    zClient.__init__c           	      C   s�   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  | j}|�� \}}}|tjkr|t� }|�|� |��  |�t� }|�|� |��  |jdk	r�|jS |jdk	r�|j�ttjd��dS )�0
        Parameters:
         - request
        �issueLiffViewNz$issueLiffView failed: unknown result)r^   �writeMessageBeginr   �CALLr_   r   r   r-   �writeMessageEndr,   �flushr]   �readMessageBegin�	EXCEPTIONr   r!   �readMessageEndrW   rX   rY   �MISSING_RESULT�	r   r   �argsr'   r(   �mtype�rseqid�x�resultr   r   r   ra   $  s*    





zClient.issueLiffViewc           	      C   s�   | j �dtj| j� t� }||_|�| j � | j ��  | j j	�
�  | j}|�� \}}}|tjkr|t� }|�|� |��  |�t� }|�|� |��  |jdk	r�|j�dS )r`   �revokeTokenN)r^   rb   r   rc   r_   rZ   r   r-   rd   r,   re   r]   rf   rg   r   r!   rh   r[   rY   rj   r   r   r   rp   A  s&    




zClient.revokeToken)N)rI   rT   rU   r   ra   rp   r   r   r   r   r\     s   
r\   )�thrift.Thriftr   r   r   r   r   �thrift.protocol.TProtocolr   �sys�logging�ttypesr	   r
   r   r   r   �thrift.transportr   �objectr   rW   rZ   r[   r\   r   r   r   r   �<module>   s   BO@@